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

def const_template(option_value, f, name):
    const_tmpl = """{% if str(o.string) != "" and o.string['const'] %}
    if p.{{ name }} != \"{{ o.string['const'] }}\":
        raise ValidationFailed(\"{{ name }} not equal to {{ o.string['const'] }}\")
    {% elif str(o.bool) != "" and o.bool['const'] != '' %}
    if p.{{ name }} != {{ o.bool['const'] }}:
        raise ValidationFailed(\"{{ name }} not equal to {{ o.bool['const'] }}\")
    {% endif %}
    """
    return Template(const_tmpl).render(o = option_value, f = f, name = name, str = str)

def in_template(value, name):
    in_tmpl = """
    {% if value['in'] %}
    if p.{{ name }} not in {{ value['in'] }}:
        raise ValidationFailed(\"{{ name }} not in {{ value['in'] }}\")
    {% endif %}
    {% if value['not_in'] %}
    if p.{{ name }} in {{ value['not_in'] }}:
        raise ValidationFailed(\"{{ name }} in {{ value['not_in'] }}\")
    {% endif %}
    """
    return Template(in_tmpl).render(value = value, name = name)

def string_template(option_value, f, name):
    str_templ = """
    {{ const_template(o, f, name) }}
    {{ in_template(o.string, name) }}
    {% set s = o.string %}
    {% if s['len'] %}
    if len(p.{{ name }}) != {{ s['len'] }}:
        raise ValidationFailed(\"{{ name }} length does not equal {{ s['len'] }}\")
    {% endif %}
    {% if s['min_len'] %}
    if len(p.{{ name }}) < {{ s['min_len'] }}:
        raise ValidationFailed(\"{{ name }} length is less than {{ s['min_len'] }}\")
    {% endif %}
    {% if s['max_len'] %}
    if len(p.{{ name }}) > {{ s['max_len'] }}:
        raise ValidationFailed(\"{{ name }} length is more than {{ s['max_len'] }}\")
    {% endif %}
    {% if s['len_bytes'] %}
    if byte_len(p.{{ name }}) != {{ s['len_bytes'] }}:
        raise ValidationFailed(\"{{ name }} length does not equal {{ s['len_bytes'] }}\")
    {% endif %}
    {% if s['min_bytes'] %}
    if byte_len(p.{{ name }}) < {{ s['min_bytes'] }}:
        raise ValidationFailed(\"{{ name }} length is less than {{ s['min_bytes'] }}\")
    {% endif %}
    {% if s['max_bytes'] %}
    if byte_len(p.{{ name }}) > {{ s['max_bytes'] }}:
        raise ValidationFailed(\"{{ name }} length is greater than {{ s['max_bytes'] }}\")
    {% endif %}
    {% if s['pattern'] %}
    if re.search(r\'{{ s['pattern'] }}\', p.{{ name }}) is None:
        raise ValidationFailed(\"{{ name }} pattern does not match {{ s['pattern'] }}\")
    {% endif %}
    {% if s['prefix'] %}
    if not p.{{ name }}.startswith(\"{{ s['prefix'] }}\"):
        raise ValidationFailed(\"{{ name }} does not start with prefix {{ s['prefix'] }}\")
    {% endif %}
    {% if s['suffix'] %}
    if not p.{{ name }}.endswith(\"{{ s['suffix'] }}\"):
        raise ValidationFailed(\"{{ name }} does not end with suffix {{ s['suffix'] }}\")
    {% endif %}
    {% if s['contains'] %}
    if not \"{{ s['contains'] }}\" in p.{{ name }}:
        raise ValidationFailed(\"{{ name }} does not contain {{ s['contains'] }}\")
    {% endif %}
    {% if s['email'] %}
    if not _validateEmail(p.{{ name }}):
        raise ValidationFailed(\"{{ name }} is not a valid email\")
    {% endif %}    
    {% if s['hostname'] %}
    if not _validateHostName(p.{{ name }}):
        raise ValidationFailed(\"{{ name }} is not a valid email\")
    {% endif %}
    {% if s['address'] %}
    try:
        ipaddress.ip_address(unicode(p.{{ name }}))
    except ValueError:
        if not _validateHostName(p.{{ name }}):
            raise ValidationFailed(\"{{ name }} is not a valid address\")
    {% endif %}
    {% if s['ip'] %}
    try:
        ipaddress.ip_address(unicode(p.{{ name }}))
    except ValueError:
        raise ValidationFailed(\"{{ name }} is not a valid ip\")
    {% endif %}
    {% if s['ipv4'] %}
    try:
        ipaddress.IPv4Address(unicode(p.{{ name }}))
    except ValueError:
        raise ValidationFailed(\"{{ name }} is not a valid ipv4\")
    {% endif %}
    {% if s['ipv6'] %}
    try:
        ipaddress.IPv6Address(unicode(p.{{ name }}))
    except ValueError:
        raise ValidationFailed(\"{{ name }} is not a valid ipv6\")
    {% endif %}
    {% if s['uri'] %}
    url = urlparse.urlparse(p.{{ name }})
    if not all([url.scheme, url.netloc, url.path]):
        raise ValidationFailed(\"{{ name }} is not a valid uri\")
    {% endif %}
    {% if s['uri_ref'] %}
    url = urlparse.urlparse(p.{{ name }})
    if not all([url.scheme, url.path]) and url.fragment:
        raise ValidationFailed(\"{{ name }} is not a valid uri ref\")
    {% endif %}
    {% if s['uuid'] %}
    try:
        uuid.UUID(p.{{ name }})
    except ValueError:
        raise ValidationFailed(\"{{ name }} is not a valid UUID\")     
    {% endif %}
    """
    return Template(str_templ).render(o=option_value,f=f,name=name,const_template=const_template, in_template=in_template)

def required_template(value, name):
    req_tmpl = """{% if value['required'] %}
    if not _has_field(p, \"{{ name }}\"):
        raise ValidationFailed(\"{{ name }} is required.\")
    {% endif %}
    """
    return Template(req_tmpl).render(value = value, name = name)

def message_template(option_value, f, name):
    message_tmpl = """{% if m.message %}
    {{ required_template(m.message, name) }}
    {% endif %}
    {% if m.message and m.message['skip'] %}
    # Skipping validation for {{ name }}
    {% else %}
    if _has_field(p, \"{{ name }}\"):
        return generate_validate(p.{{name}})(p.{{name}})    
    {% endif %}
    """
    return Template(message_tmpl).render(m=option_value,f=f, name=name, generate_validate=generate_validate, required_template=required_template)

def bool_template(option_value, f, name):
    bool_tmpl = """
    {{ const_template(o, f, name) }}
    """
    return Template(bool_tmpl).render(o=option_value,f=f,name=name,const_template=const_template)

def num_template(option_value, f, name, num):
    num_tmpl = """
    {% if num.HasField('const') and str(o.float) == "" %}
    if p.{{ name }} != {{ num['const'] }}:
        raise ValidationFailed(\"{{ name }} not equal to {{ num['const'] }}\")
    {% endif %}
    {% if num.HasField('const') and str(o.float) != "" %}
    if p.{{ name }} != struct.unpack(\"f\", struct.pack(\"f\", ({{ num['const'] }})))[0]:
        raise ValidationFailed(\"{{ name }} not equal to {{ num['const'] }}\")
    {% endif %}
    {{ in_template(num, name) }}
    {% if num.HasField('lt') %}
        {% if num.HasField('gt') %}
            {% if num['lt'] > num['gt'] %}
    if p.{{ name }} <= {{ num['gt'] }} or p.{{ name }} >= {{ num ['lt'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['lt'], num['gt'] }}\")
            {% else %}
    if p.{{ name }} >= {{ num['lt'] }} and p.{{ name }} <= {{ num['gt'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['gt'], num['lt'] }}\")
            {% endif %}
        {% elif num.HasField('gte') %}
            {% if num['lt'] > num['gte'] %}
    if p.{{ name }} < {{ num['gte'] }} or p.{{ name }} >= {{ num ['lt'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['lt'], num['gte'] }}\")
            {% else %}
    if p.{{ name }} >= {{ num['lt'] }} and p.{{ name }} < {{ num['gte'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['gte'], num['lt'] }}\")
            {% endif %}
        {% else %}
    if p.{{ name }} >= {{ num['lt'] }}:
        raise ValidationFailed(\"{{ name }} is not lesser than {{ num['lt'] }}\")
        {% endif %}
    {% elif num.HasField('lte') %}
        {% if num.HasField('gt') %}
            {% if num['lte'] > num['gt'] %}
    if p.{{ name }} <= {{ num['gt'] }} or p.{{ name }} > {{ num ['lte'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['lte'], num['gt'] }}\")
            {% else %}
    if p.{{ name }} > {{ num['lte'] }} and p.{{ name }} <= {{ num['gt'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['gt'], num['lte'] }}\")
            {% endif %}
        {% elif num.HasField('gte') %}
            {% if num['lte'] > num['gte'] %}
    if p.{{ name }} < {{ num['gte'] }} or p.{{ name }} > {{ num ['lte'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['lte'], num['gte'] }}\")
            {% else %}
    if p.{{ name }} > {{ num['lte'] }} and p.{{ name }} < {{ num['gte'] }}:
        raise ValidationFailed(\"{{ name }} is not in range {{ num['gte'], num['lte'] }}\")
            {% endif %}
        {% else %}
    if p.{{ name }} > {{ num['lte'] }}:
        raise ValidationFailed(\"{{ name }} is not lesser than or equal to {{ num['lte'] }}\")
        {% endif %}
    {% elif num.HasField('gt') %}
    if p.{{ name }} <= {{ num['gt'] }}:
        raise ValidationFailed(\"{{ name }} is not greater than {{ num['gt'] }}\")
    {% elif num.HasField('gte') %}
    if p.{{ name }} < {{ num['gte'] }}:
        raise ValidationFailed(\"{{ name }} is not greater than or equal to {{ num['gte'] }}\")    
    {% endif %}
    
    """
    return Template(num_tmpl).render(o=option_value,f=f,name=name,num=num,in_template=in_template,str=str)

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

def duration_template(option_value, f, name):
    dur_tmpl = """
    {{ required_template(o.duration, name) }}
    if _has_field(p, \"{{ name }}\"):
        dur = p.{{ name }}.seconds + round((10**-9 * p.{{ name }}.nanos),9)
        {% set dur = o.duration %}
        {% if dur.HasField('const') %}
        if dur != {{ dur_lit(dur['const']) }}:
            raise ValidationFailed(\"{{ name }} is not equal to {{ dur_lit(dur['const']) }}\")
        {% endif %}
        {% if dur['in'] %}
        if dur not in {{ dur_arr(dur['in']) }}:
            raise ValidationFailed(\"{{ name }} is not in {{ dur_arr(dur['in']) }}\") 
        {% endif %}
        {% if dur['not_in'] %}
        if dur in {{ dur_arr(dur['not_in']) }}:
            raise ValidationFailed(\"{{ name }} is not in {{ dur_arr(dur['not_in']) }}\") 
        {% endif %}
        {% if dur.HasField('lt') %}
            {% if dur.HasField('gt') %}
                {% if dur_lit(dur['lt']) > dur_lit(dur['gt']) %}
        if dur <= {{ dur_lit(dur['gt']) }} or dur >= {{ dur_lit(dur['lt']) }}:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['lt']), dur_lit(dur['gt']) }}\")
                {% else %}
        if dur >= {{ dur_lit(dur['lt']) }} and dur <= {{ dur_lit(dur['gt']) }}:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['gt']), dur_lit(dur['lt']) }}\")
                {% endif %}
            {% elif dur.HasField('gte') %}
                {% if dur_lit(dur['lt']) > dur_lit(dur['gte']) %}
        if dur < {{ dur_lit(dur['gte']) }} or dur >= {{ dur_lit(dur['lt']) }}:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['lt']), dur_lit(dur['gte']) }}\")
                {% else %}
        if dur >= {{ dur_lit(dur['lt']) }} and dur < {{ dur_lit(dur['gte']) }}:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['gte']), dur_lit(dur['lt']) }}\")
                {% endif %}
            {% else %}
        if dur >= {{ dur_lit(dur['lt']) }}:
            raise ValidationFailed(\"{{ name }} is not lesser than {{ dur_lit(dur['lt']) }}\")    
            {% endif %}
        {% elif dur.HasField('lte') %}
            {% if dur.HasField('gt') %}
                {% if dur_lit(dur['lte']) > dur_lit(dur['gt']) %}
        if dur <= {{ dur_lit(dur['gt']) }} or dur > {{ dur_lit(dur['lte']) }}:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['lte']), dur_lit(dur['gt']) }}\")
                {% else %}
        if dur > {{ dur_lit(dur['lte']) }} and dur <= {{ dur_lit(dur['gt']) }}:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['gt']), dur_lit(dur['lte']) }}\")
                {% endif %}
            {% elif dur.HasField('gte') %}
                {% if dur_lit(dur['lte']) > dur_lit(dur['gte']) %}
        if dur < {{ dur_lit(dur['gte']) }} or dur > {{ dur_lit(dur['lte']) }}:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['lte']), dur_lit(dur['gte']) }}\")
                {% else %}
        if dur > {{ dur_lit(dur['lte']) }} and dur < {{ dur_lit(dur['gte']) }}:
            raise ValidationFailed(\"{{ name }} is not in range {{ dur_lit(dur['gte']), dur_lit(dur['lte']) }}\")
                {% endif %}
            {% else %}
        if dur > {{ dur_lit(dur['lte']) }}:
            raise ValidationFailed(\"{{ name }} is not lesser than or equal to {{ dur_lit(dur['lte']) }}\")                   
            {% endif %}
        {% elif dur.HasField('gt') %}
        if dur <= {{ dur_lit(dur['gt']) }}:
            raise ValidationFailed(\"{{ name }} is not greater than {{ dur_lit(dur['gt']) }}\")
        {% elif dur.HasField('gte') %}
        if dur < {{ dur_lit(dur['gte']) }}:
            raise ValidationFailed(\"{{ name }} is not greater than or equal to {{ dur_lit(dur['gte']) }}\")
        {% endif %}
    """
    return Template(dur_tmpl).render(o=option_value,f=f,name=name,required_template=required_template, _has_field=_has_field,dur_lit=dur_lit,dur_arr=dur_arr)

def rule_type(field, name = ""):
    if has_validate(field) and field.message_type is None and not field.containing_oneof:
        for option_descriptor, option_value in field.GetOptions().ListFields():
            if option_descriptor.full_name == "validate.rules":
                if str(option_value.string) is not "":
                    return string_template(option_value, field, ".".join([x for x in [name, field.name] if x]))
                elif str(option_value.message) is not "":
                    return message_template(option_value, field, ".".join([x for x in [name, field.name] if x]))
                elif str(option_value.bool) is not "":
                    return bool_template(option_value, field, ".".join([x for x in [name, field.name] if x]))
                elif str(option_value.float) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.float)
                elif str(option_value.double) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.double)
                elif str(option_value.int32) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.int32)
                elif str(option_value.int64) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.int64)
                elif str(option_value.uint32) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.uint32)
                elif str(option_value.uint64) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.uint64)
                elif str(option_value.sint32) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.sint32)
                elif str(option_value.sint64) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.sint64)
                elif str(option_value.fixed32) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.fixed32)
                elif str(option_value.fixed64) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.fixed64)
                elif str(option_value.sfixed32) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.sfixed32)
                elif str(option_value.sfixed64) is not "":
                    return num_template(option_value, field, ".".join([x for x in [name, field.name] if x]), option_value.sfixed64)
                else:
                    return "raise UnimplementedException()"
    if field.message_type and not field.containing_oneof:
        for option_descriptor, option_value in field.GetOptions().ListFields():
            if option_descriptor.full_name == "validate.rules":
                if str(option_value.duration) is not "":
                    return duration_template(option_value, field, ".".join([x for x in [name, field.name] if x]))
                else:
                    return message_template(option_value, field, ".".join([x for x in [name, field.name] if x]))
        if field.message_type.full_name.startswith("google.protobuf"):
            return ""
        else:
            return message_template(None, field, ".".join([x for x in [name, field.name] if x]))


def file_template(proto_message):
    file_tmp = """def validate(p):
    {% set accessor = p.DESCRIPTOR -%}
    {% for option_descriptor, option_value in accessor.GetOptions().ListFields() %}
        {% if option_descriptor.full_name == "validate.disabled" and option_value %}
    return None
        {% endif %}
    {% endfor %}
    {% for field in accessor.fields -%}
        {{ rule_type(field) }}
    {%- endfor %}
    return None"""
    return Template(file_tmp).render(rule_type=rule_type, p=proto_message, dir=dir)


class UnimplementedException(Exception):
    pass

class ValidationFailed(Exception):
    pass

def generate_validate(proto_message):
    func = file_template(proto_message)
    exec(func); return validate
