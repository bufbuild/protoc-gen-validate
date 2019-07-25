import re
from validate_email import validate_email
import ipaddress
try:
    import urlparse
except ImportError:
    import urllib.parse as urlparse
import uuid
import struct
from jinja2 import Template
import time

printer = ""

def generate_validate(proto_message):
    func = file_template(proto_message)
    global printer
    printer += func + "\n"
    exec(func); return validate

def print_validate(proto_message):
    return printer

def has_validate(field):
    if field.GetOptions() is None:
        return False
    for option_descriptor, option_value in field.GetOptions().ListFields():
        if option_descriptor.full_name == "validate.rules":
            return True
    return False

def byte_len(s):
    try:
        return len(s.encode('utf-8'))
    except:
        return len(s)

def _validateHostName(host):
    if len(host) > 253:
        return False

    s = host.rsplit(".",1)[0].lower()

    for part in s.split("."):
        if len(part) == 0 or len(part) > 63:
            return False

        # Host names cannot begin or end with hyphens
        if s[0] == "-" or s[len(s)-1] == '-':
            return False
        for r in part:
            if (r < 'A' or r > 'Z') and (r < 'a' or r > 'z') and (r < '0' or r > '9') and r != '-':
                return False
    return True


def _validateEmail(addr):
    if '<' in addr and '>' in addr: addr = addr.split("<")[1].split(">")[0]

    if not validate_email(addr):
        return False

    if len(addr) > 254:
        return False

    parts = addr.split("@")
    if len(parts[0]) > 64:
        return False
    return _validateHostName(parts[1])

def _has_field(message_pb, property_name):
    # NOTE: As of proto3, HasField() only works for message fields, not for
    #       singular (non-message) fields. First try to use HasField and
    #       if it fails (with a ValueError) we manually consult the fields.
    try:
        return message_pb.HasField(property_name)
    except:
        all_fields = set([field.name for field in message_pb.DESCRIPTOR.fields])
        return property_name in all_fields

def const_template(option_value, name):
    const_tmpl = """{%- if str(o.string) and o.string.HasField('const') -%}
    if p.{{ name }} != \"{{ o.string['const'] }}\":
        raise ValidationFailed(\"{{ name }} not equal to {{ o.string['const'] }}\")
    {%- elif str(o.bool) and o.bool['const'] != "" -%}
    if p.{{ name }} != {{ o.bool['const'] }}:
        raise ValidationFailed(\"{{ name }} not equal to {{ o.bool['const'] }}\")
    {%- endif -%}
    """
    return Template(const_tmpl).render(o = option_value, name = name, str = str)

def in_template(value, name):
    in_tmpl = """
    {%- if value['in'] %}
    if p.{{ name }} not in {{ value['in'] }}:
        raise ValidationFailed(\"{{ name }} not in {{ value['in'] }}\")
    {%- endif -%}
    {%- if value['not_in'] %}
    if p.{{ name }} in {{ value['not_in'] }}:
        raise ValidationFailed(\"{{ name }} in {{ value['not_in'] }}\")
    {%- endif -%}
    """
    return Template(in_tmpl).render(value = value, name = name)

def string_template(option_value, name):
    str_templ = """
    {{ const_template(o, name) -}}
    {{ in_template(o.string, name) -}}
    {%- set s = o.string -%}
    {%- if s['len'] %}
    if len(p.{{ name }}) != {{ s['len'] }}:
        raise ValidationFailed(\"{{ name }} length does not equal {{ s['len'] }}\")
    {%- endif -%}
    {%- if s['min_len'] %}
    if len(p.{{ name }}) < {{ s['min_len'] }}:
        raise ValidationFailed(\"{{ name }} length is less than {{ s['min_len'] }}\")
    {%- endif -%}
    {%- if s['max_len'] %}
    if len(p.{{ name }}) > {{ s['max_len'] }}:
        raise ValidationFailed(\"{{ name }} length is more than {{ s['max_len'] }}\")
    {%- endif -%}
    {%- if s['len_bytes'] %}
    if byte_len(p.{{ name }}) != {{ s['len_bytes'] }}:
        raise ValidationFailed(\"{{ name }} length does not equal {{ s['len_bytes'] }}\")
    {%- endif -%}
    {%- if s['min_bytes'] %}
    if byte_len(p.{{ name }}) < {{ s['min_bytes'] }}:
        raise ValidationFailed(\"{{ name }} length is less than {{ s['min_bytes'] }}\")
    {%- endif -%}
    {%- if s['max_bytes'] %}
    if byte_len(p.{{ name }}) > {{ s['max_bytes'] }}:
        raise ValidationFailed(\"{{ name }} length is greater than {{ s['max_bytes'] }}\")
    {%- endif -%}
    {%- if s['pattern'] %}
    if re.search(r\'{{ s['pattern'] }}\', p.{{ name }}) is None:
        raise ValidationFailed(\"{{ name }} pattern does not match {{ s['pattern'] }}\")
    {%- endif -%}
    {%- if s['prefix'] %}
    if not p.{{ name }}.startswith(\"{{ s['prefix'] }}\"):
        raise ValidationFailed(\"{{ name }} does not start with prefix {{ s['prefix'] }}\")
    {%- endif -%}
    {%- if s['suffix'] %}
    if not p.{{ name }}.endswith(\"{{ s['suffix'] }}\"):
        raise ValidationFailed(\"{{ name }} does not end with suffix {{ s['suffix'] }}\")
    {%- endif -%}
    {%- if s['contains'] %}
    if not \"{{ s['contains'] }}\" in p.{{ name }}:
        raise ValidationFailed(\"{{ name }} does not contain {{ s['contains'] }}\")
    {%- endif -%}
    {%- if s['email'] %}
    if not _validateEmail(p.{{ name }}):
        raise ValidationFailed(\"{{ name }} is not a valid email\")
    {%- endif -%}    
    {%- if s['hostname'] %}
    if not _validateHostName(p.{{ name }}):
        raise ValidationFailed(\"{{ name }} is not a valid email\")
    {%- endif -%}
    {%- if s['address'] %}
    try:
        ipaddress.ip_address(unicode(p.{{ name }}))
    except ValueError:
        if not _validateHostName(p.{{ name }}):
            raise ValidationFailed(\"{{ name }} is not a valid address\")
    {%- endif -%}
    {%- if s['ip'] %}
    try:
        ipaddress.ip_address(unicode(p.{{ name }}))
    except ValueError:
        raise ValidationFailed(\"{{ name }} is not a valid ip\")
    {%- endif -%}
    {%- if s['ipv4'] %}
    try:
        ipaddress.IPv4Address(unicode(p.{{ name }}))
    except ValueError:
        raise ValidationFailed(\"{{ name }} is not a valid ipv4\")
    {%- endif -%}
    {%- if s['ipv6'] %}
    try:
        ipaddress.IPv6Address(unicode(p.{{ name }}))
    except ValueError:
        raise ValidationFailed(\"{{ name }} is not a valid ipv6\")
    {%- endif %}
    {%- if s['uri'] %}
    url = urlparse.urlparse(p.{{ name }})
    if not all([url.scheme, url.netloc, url.path]):
        raise ValidationFailed(\"{{ name }} is not a valid uri\")
    {%- endif %}
    {%- if s['uri_ref'] %}
    url = urlparse.urlparse(p.{{ name }})
    if not all([url.scheme, url.path]) and url.fragment:
        raise ValidationFailed(\"{{ name }} is not a valid uri ref\")
    {%- endif -%}
    {%- if s['uuid'] %}
    try:
        uuid.UUID(p.{{ name }})
    except ValueError:
        raise ValidationFailed(\"{{ name }} is not a valid UUID\")     
    {%- endif -%}
    """
    return Template(str_templ).render(o = option_value, name = name, const_template = const_template, in_template = in_template)

def required_template(value, name):
    req_tmpl = """{%- if value['required'] -%}
    if not _has_field(p, \"{{ name }}\"):
        raise ValidationFailed(\"{{ name }} is required.\")
    {%- endif -%}
    """
    return Template(req_tmpl).render(value = value, name = name)

def message_template(option_value, name):
    message_tmpl = """{%- if m.message %}
    {{- required_template(m.message, name) }}
    {%- endif -%}
    {%- if m.message and m.message['skip'] %}
    # Skipping validation for {{ name }}
    {%- else %}
    if _has_field(p, \"{{ name }}\"):
        embedded = generate_validate(p.{{ name }})(p.{{ name }})
        if embedded is not None:
            return embedded
    {%- endif -%}
    """
    return Template(message_tmpl).render(m = option_value, name = name, required_template = required_template)

def bool_template(option_value, name):
    bool_tmpl = """
    {{ const_template(o, name) -}}
    """
    return Template(bool_tmpl).render(o = option_value, name = name, const_template = const_template)

def num_template(option_value, name, num):
    num_tmpl = """{%- if num.HasField('const') and str(o.float) == "" -%}
    if p.{{ name }} != {{ num['const'] }}:
        raise ValidationFailed(\"{{ name }} not equal to {{ num['const'] }}\")
    {%- endif -%}
    {%- if num.HasField('const') and str(o.float) != "" %}
    if p.{{ name }} != struct.unpack(\"f\", struct.pack(\"f\", ({{ num['const'] }})))[0]:
        raise ValidationFailed(\"{{ name }} not equal to {{ num['const'] }}\")
    {%- endif -%}
    {{ in_template(num, name) }}
    {%- if num.HasField('lt') %}
        {%- if num.HasField('gt') %}
            {%- if num['lt'] > num['gt'] %}
    if p.{{ name }} <= {{ num['gt'] }} or p.{{ name }} >= {{ num ['lt'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['lt'], num['gt'] }}\")
            {%- else %}
    if p.{{ name }} >= {{ num['lt'] }} and p.{{ name }} <= {{ num['gt'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['gt'], num['lt'] }}\")
            {%- endif -%}
        {%- elif num.HasField('gte') %}
            {%- if num['lt'] > num['gte'] %}
    if p.{{ name }} < {{ num['gte'] }} or p.{{ name }} >= {{ num ['lt'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['lt'], num['gte'] }}\")
            {%- else %}
    if p.{{ name }} >= {{ num['lt'] }} and p.{{ name }} < {{ num['gte'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['gte'], num['lt'] }}\")
            {%- endif -%}
        {%- else %}
    if p.{{ name }} >= {{ num['lt'] }}:
        raise ValidationFailed(\"{{ name }} is not lesser than {{ num['lt'] }}\")
        {%- endif -%}
    {%- elif num.HasField('lte') %}
        {%- if num.HasField('gt') %}
            {%- if num['lte'] > num['gt'] %}
    if p.{{ name }} <= {{ num['gt'] }} or p.{{ name }} > {{ num ['lte'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['lte'], num['gt'] }}\")
            {%- else %}
    if p.{{ name }} > {{ num['lte'] }} and p.{{ name }} <= {{ num['gt'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['gt'], num['lte'] }}\")
            {%- endif -%}
        {%- elif num.HasField('gte') %}
            {%- if num['lte'] > num['gte'] %}
    if p.{{ name }} < {{ num['gte'] }} or p.{{ name }} > {{ num ['lte'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['lte'], num['gte'] }}\")
            {%- else %}
    if p.{{ name }} > {{ num['lte'] }} and p.{{ name }} < {{ num['gte'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['gte'], num['lte'] }}\")
            {%- endif -%}
        {%- else %}
    if p.{{ name }} > {{ num['lte'] }}:
        raise ValidationFailed(\"{{ name }} is not lesser than or equal to {{ num['lte'] }}\")
        {%- endif -%}
    {%- elif num.HasField('gt') %}
    if p.{{ name }} <= {{ num['gt'] }}:
        raise ValidationFailed(\"{{ name }} is not greater than {{ num['gt'] }}\")
    {%- elif num.HasField('gte') %}
    if p.{{ name }} < {{ num['gte'] }}:
        raise ValidationFailed(\"{{ name }} is not greater than or equal to {{ num['gte'] }}\")    
    {%- endif -%}
    """
    return Template(num_tmpl).render(o = option_value, name = name, num = num, in_template = in_template, str = str)

def dur_arr(dur):
    value = 0
    arr = []
    for val in dur:
        value += val.seconds
        value += (10**-9 * val.nanos)
        arr.append(value)
        value = 0
    return arr

def dur_lit(dur):
    value = dur.seconds + (10**-9 * dur.nanos)
    return value

def duration_template(option_value, name):
    dur_tmpl = """
    {{- required_template(o.duration, name) }}
    if _has_field(p, \"{{ name }}\"):
        dur = p.{{ name }}.seconds + round((10**-9 * p.{{ name }}.nanos), 9)
        {%- set dur = o.duration -%}
        {%- if dur.HasField('lt') %} 
        lt = {{ dur_lit(dur['lt']) }} 
        {% endif %}
        {%- if dur.HasField('lte') %} 
        lte = {{ dur_lit(dur['lte']) }} 
        {% endif %}
        {%- if dur.HasField('gt') %} 
        gt = {{ dur_lit(dur['gt']) }} 
        {% endif %}
        {%- if dur.HasField('gte') %} 
        gte = {{ dur_lit(dur['gte']) }} 
        {% endif %}
        {%- if dur.HasField('const') %}
        if dur != {{ dur_lit(dur['const']) }}:
            raise ValidationFailed(\"{{ name }} is not equal to {{ dur_lit(dur['const']) }}\")
        {%- endif -%}
        {%- if dur['in'] %}
        if dur not in {{ dur_arr(dur['in']) }}:
            raise ValidationFailed(\"{{ name }} is not in {{ dur_arr(dur['in']) }}\") 
        {%- endif -%}
        {%- if dur['not_in'] %}
        if dur in {{ dur_arr(dur['not_in']) }}:
            raise ValidationFailed(\"{{ name }} is not in {{ dur_arr(dur['not_in']) }}\") 
        {%- endif -%}
        {%- if dur.HasField('lt') %}
            {%- if dur.HasField('gt') %}
                {%- if dur_lit(dur['lt']) > dur_lit(dur['gt']) %}
        if dur <= gt or dur >= lt:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['lt']), dur_lit(dur['gt']) }}\")
                {%- else -%}
        if dur >= lt and dur <= gt:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['gt']), dur_lit(dur['lt']) }}\")
                {%- endif -%}
            {%- elif dur.HasField('gte') %}
                {%- if dur_lit(dur['lt']) > dur_lit(dur['gte']) %}
        if dur < gte or dur >= lt:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['lt']), dur_lit(dur['gte']) }}\")
                {%- else -%}
        if dur >= lt and dur < gte:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['gte']), dur_lit(dur['lt']) }}\")
                {%- endif -%}
            {%- else -%}
        if dur >= lt:
            raise ValidationFailed(\"{{ name }} is not lesser than {{ dur_lit(dur['lt']) }}\")    
            {%- endif -%}
        {%- elif dur.HasField('lte') %}
            {%- if dur.HasField('gt') %}
                {%- if dur_lit(dur['lte']) > dur_lit(dur['gt']) %}
        if dur <= gt or dur > lte:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['lte']), dur_lit(dur['gt']) }}\")
                {%- else -%}
        if dur > lte and dur <= gt:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['gt']), dur_lit(dur['lte']) }}\")
                {%- endif -%}
            {%- elif dur.HasField('gte') %}
                {%- if dur_lit(dur['lte']) > dur_lit(dur['gte']) %}
        if dur < gte or dur > lte:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['lte']), dur_lit(dur['gte']) }}\")
                {%- else -%}
        if dur > lte and dur < gte:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['gte']), dur_lit(dur['lte']) }}\")
                {%- endif -%}
            {%- else -%}
        if dur > lte:
            raise ValidationFailed(\"{{ name }} is not lesser than or equal to {{ dur_lit(dur['lte']) }}\")                   
            {%- endif -%}
        {%- elif dur.HasField('gt') %}
        if dur <= gt:
            raise ValidationFailed(\"{{ name }} is not greater than {{ dur_lit(dur['gt']) }}\")
        {%- elif dur.HasField('gte') %}
        if dur < gte:
            raise ValidationFailed(\"{{ name }} is not greater than or equal to {{ dur_lit(dur['gte']) }}\")
        {%- endif -%}
    """
    return Template(dur_tmpl).render(o = option_value, name = name, required_template = required_template, dur_lit = dur_lit, dur_arr = dur_arr)

def timestamp_template(option_value, name):
    timestamp_tmpl = """
    {{- required_template(o.timestamp, name) }}
    if _has_field(p, \"{{ name }}\"):
        ts = p.{{ name }}.seconds + round((10**-9 * p.{{ name }}.nanos), 9)
        {%- set ts = o.timestamp -%}
        {%- if ts.HasField('lt') %} 
        lt = {{ dur_lit(ts['lt']) }} 
        {% endif -%}
        {%- if ts.HasField('lte') %} 
        lte = {{ dur_lit(ts['lte']) }} 
        {% endif -%}
        {%- if ts.HasField('gt') %} 
        gt = {{ dur_lit(ts['gt']) }} 
        {% endif -%}
        {%- if ts.HasField('gte') %} 
        gte = {{ dur_lit(ts['gte']) }} 
        {% endif -%}
        {%- if ts.HasField('const') %}
        if ts != {{ dur_lit(ts['const']) }}:
            raise ValidationFailed(\"{{ name }} is not equal to {{ dur_lit(ts['const']) }}\")
        {% endif %}
        {%- if ts['in'] %}
        if ts not in {{ dur_arr(ts['in']) }}:
            raise ValidationFailed(\"{{ name }} is not in {{ dur_arr(ts['in']) }}\") 
        {%- endif %}
        {%- if ts['not_in'] %}
        if ts in {{ dur_arr(ts['not_in']) }}:
            raise ValidationFailed(\"{{ name }} is not in {{ dur_arr(ts['not_in']) }}\") 
        {%- endif %}
        {%- if ts.HasField('lt') %}
            {%- if ts.HasField('gt') %}
                {%- if dur_lit(ts['lt']) > dur_lit(ts['gt']) %}
        if ts <= gt or ts >= lt:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(ts['lt']), dur_lit(ts['gt']) }}\")
                {%- else -%}
        if ts >= lt and ts <= gt:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(ts['gt']), dur_lit(ts['lt']) }}\")
                {%- endif -%}
            {%- elif ts.HasField('gte') %}
                {%- if dur_lit(ts['lt']) > dur_lit(ts['gte']) %}
        if ts < gte or ts >= lt:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(ts['lt']), dur_lit(ts['gte']) }}\")
                {%- else -%}
        if ts >= lt and ts < gte:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(ts['gte']), dur_lit(ts['lt']) }}\")
                {%- endif -%}
            {%- else -%}
        if ts >= lt:
            raise ValidationFailed(\"{{ name }} is not lesser than {{ dur_lit(ts['lt']) }}\")    
            {%- endif -%}
        {%- elif ts.HasField('lte') %}
            {%- if ts.HasField('gt') %}
                {%- if dur_lit(ts['lte']) > dur_lit(ts['gt']) %}
        if ts <= gt or ts > lte:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(ts['lte']), dur_lit(ts['gt']) }}\")
                {%- else -%}
        if ts > lte and ts <= gt:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(ts['gt']), dur_lit(ts['lte']) }}\")
                {%- endif -%}
            {%- elif ts.HasField('gte') %}
                {%- if dur_lit(ts['lte']) > dur_lit(ts['gte']) %}
        if ts < gte or ts > lte:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(ts['lte']), dur_lit(ts['gte']) }}\")
                {%- else -%}
        if ts > lte and ts < gte:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(ts['gte']), dur_lit(ts['lte']) }}\")
                {%- endif -%}
            {%- else -%}
        if ts > lte:
            raise ValidationFailed(\"{{ name }} is not lesser than or equal to {{ dur_lit(ts['lte']) }}\")                   
            {%- endif -%}
        {%- elif ts.HasField('gt') %}
        if ts <= gt:
            raise ValidationFailed(\"{{ name }} is not greater than {{ dur_lit(ts['gt']) }}\")
        {%- elif ts.HasField('gte') %}
        if ts < gte:
            raise ValidationFailed(\"{{ name }} is not greater than or equal to {{ dur_lit(ts['gte']) }}\")
        {%- elif ts.HasField('lt_now') %}
        now = time.time()
            {%- if ts.HasField('within') %}
        within = {{ dur_lit(ts['within']) }}
        if ts >= now or ts >= now - within:
            raise ValidationFailed(\"{{ name }} is not within range {{ dur_lit(ts['within']) }}\")
            {%- else %}
        if ts >= now:
            raise ValidationFailed(\"{{ name }} is not lesser than now\")
            {%- endif -%} 
        {%- elif ts.HasField('gt_now') %}
        now = time.time()
            {%- if ts.HasField('within') %}
        within = {{ dur_lit(ts['within']) }}
        if ts <= now or ts <= now + within:
            raise ValidationFailed(\"{{ name }} is not within range {{ dur_lit(ts['within']) }}\")
            {%- else %}
        if ts <= now:
            raise ValidationFailed(\"{{ name }} is not greater than now\")    
            {%- endif -%}
        {%- elif ts.HasField('within') %}
        now = time.time()
        within = {{ dur_lit(ts['within']) }}
        if ts >= now + within or ts <= now - within:
             raise ValidationFailed(\"{{ name }} is not within range {{ dur_lit(ts['within']) }}\")
        {%- endif -%}
    """
    return Template(timestamp_tmpl).render(o = option_value, name = name, required_template = required_template, dur_lit = dur_lit, dur_arr = dur_arr)

def wrapper_template(option_value, name):
    wrapper_tmpl = """
    if p.HasField(\"{{ name }}\"):
        {%- if str(option_value.float) %}
        {{- num_template(option_value, name + ".value", option_value.float)|indent(8,True) -}}
        {% endif -%}
        {%- if str(option_value.double) %}
        {{- num_template(option_value, name + ".value", option_value.double)|indent(8,True) -}}
        {% endif -%}
        {%- if str(option_value.int32) %}
        {{- num_template(option_value, name + ".value", option_value.int32)|indent(8,True) -}}
        {% endif -%}
        {%- if str(option_value.int64) %}
        {{- num_template(option_value, name + ".value", option_value.int64)|indent(8,True) -}}
        {% endif -%}
        {%- if str(option_value.uint32) %}
        {{- num_template(option_value, name + ".value", option_value.uint32)|indent(8,True) -}}
        {% endif -%}
        {%- if str(option_value.uint64) %}
        {{- num_template(option_value, name + ".value", option_value.uint64)|indent(8,True) -}}
        {% endif -%}
        {%- if str(option_value.bool) %}
        {{- bool_template(option_value, name + ".value")|indent(8,True) -}}
        {% endif -%}
        {%- if str(option_value.string) %}
        {{- string_template(option_value, name + ".value")|indent(8,True) -}}
        {% endif -%}
    {%- if str(option_value.message) and option_value.message['required'] %}
    else:
        raise ValidationFailed(\"{{ name }} is required.\")
    {%- endif %}
    """
    return Template(wrapper_tmpl).render(option_value = option_value, name = name, str = str, num_template = num_template, bool_template = bool_template, string_template = string_template)

def rule_type(field):
    if has_validate(field) and field.message_type is None and not field.containing_oneof:
        for option_descriptor, option_value in field.GetOptions().ListFields():
            if option_descriptor.full_name == "validate.rules":
                if str(option_value.string):
                    return string_template(option_value, field.name)
                elif str(option_value.message):
                    return message_template(option_value, field.name)
                elif str(option_value.bool):
                    return bool_template(option_value, field.name)
                elif str(option_value.float):
                    return num_template(option_value, field.name, option_value.float)
                elif str(option_value.double):
                    return num_template(option_value, field.name, option_value.double)
                elif str(option_value.int32):
                    return num_template(option_value, field.name, option_value.int32)
                elif str(option_value.int64):
                    return num_template(option_value, field.name, option_value.int64)
                elif str(option_value.uint32):
                    return num_template(option_value, field.name, option_value.uint32)
                elif str(option_value.uint64):
                    return num_template(option_value, field.name, option_value.uint64)
                elif str(option_value.sint32):
                    return num_template(option_value, field.name, option_value.sint32)
                elif str(option_value.sint64):
                    return num_template(option_value, field.name, option_value.sint64)
                elif str(option_value.fixed32):
                    return num_template(option_value, field.name, option_value.fixed32)
                elif str(option_value.fixed64):
                    return num_template(option_value, field.name, option_value.fixed64)
                elif str(option_value.sfixed32):
                    return num_template(option_value, field.name, option_value.sfixed32)
                elif str(option_value.sfixed64):
                    return num_template(option_value, field.name, option_value.sfixed64)
                else:
                    return "raise UnimplementedException()"
    if field.message_type and not field.containing_oneof:
        for option_descriptor, option_value in field.GetOptions().ListFields():
            if option_descriptor.full_name == "validate.rules":
                if str(option_value.duration):
                    return duration_template(option_value, field.name)
                elif str(option_value.timestamp):
                    return timestamp_template(option_value, field.name)
                elif str(option_value.float) or str(option_value.int32) or str(option_value.int64) or \
                        str(option_value.double) or str(option_value.uint32) or str(option_value.uint64) or \
                        str(option_value.bool) or str(option_value.string):
                    return wrapper_template(option_value, field.name)
                elif str(option_value.bytes):
                    return "raise UnimplementedException()"
                elif str(option_value.message) is not "":
                    return message_template(option_value, field.name)
                else:
                    return "raise UnimplementedException()"
        if field.message_type.full_name.startswith("google.protobuf"):
            return ""
        else:
            return message_template(None, field.name)
    return ""

def file_template(proto_message):
    file_tmp = """
# Validates {{ p.DESCRIPTOR.name }}
def validate(p):
    {%- for option_descriptor, option_value in p.DESCRIPTOR.GetOptions().ListFields() %}
        {%- if option_descriptor.full_name == "validate.disabled" and option_value %}
    return None
        {%- endif -%}
    {%- endfor -%}
    {%- for field in p.DESCRIPTOR.fields -%}
        {%- if field.label == 3 or field.containing_oneof %}
    raise UnimplementedException()
        {%- else %}
    {{ rule_type(field) -}}
        {%- endif -%}
    {%- endfor %}
    return None"""
    return Template(file_tmp).render(rule_type = rule_type, p = proto_message)

class UnimplementedException(Exception):
    pass

class ValidationFailed(Exception):
    pass
