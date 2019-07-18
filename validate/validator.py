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

# def num_template(option_value, f, num):
#     returnval = ""
#     if str(option_value.float) is not "": # Need to convert double to float
#         if num.HasField('const'):
#             returnval += '\n  if struct.unpack(\"f\", struct.pack(\"f\", %s))[0] != p.%s:\n    return False' %(getattr(num,'const'), f.name)
#     else:
#         if num.HasField('const'):
#             returnval += '\n  if %s != p.%s:\n    return False' %(getattr(num,'const'), f.name)
#     if num.HasField('lt'):
#         if num.HasField('gt'):
#             if getattr(num, 'lt') > getattr(num, 'gt'):
#                 returnval += '\n  if p.%s <= %s or p.%s >= %s:\n    return False' %(f.name, getattr(num,'gt'), f.name, getattr(num,'lt'))
#             else:
#                 returnval += '\n  if p.%s >= %s and p.%s <= %s:\n    return False' %(f.name, getattr(num,'lt'), f.name, getattr(num,'gt'))
#         elif num.HasField('gte'):
#             if getattr(num, 'lt') > getattr(num, 'gte'):
#                 returnval += '\n  if p.%s < %s or p.%s >= %s:\n    return False' %(f.name, getattr(num,'gte'), f.name, getattr(num,'lt'))
#             else:
#                 returnval += '\n  if p.%s >= %s and p.%s < %s:\n    return False' %(f.name, getattr(num,'lt'), f.name, getattr(num,'gte'))
#         else:
#             returnval += '\n  if p.%s >= struct.unpack(\"f\", struct.pack(\"f\", %s))[0]:\n    return False' %(f.name, getattr(num,'lt'))
#     elif num.HasField('lte'):
#         if num.HasField('gt'):
#             if getattr(num, 'lte') > getattr(num, 'gt'):
#                 returnval += '\n  if p.%s <= %s or p.%s > %s:\n    return False' %(f.name, getattr(num,'gt'), f.name, getattr(num,'lte'))
#             else:
#                 returnval += '\n  if p.%s > %s and p.%s <= %s:\n    return False' %(f.name, getattr(num,'lte'), f.name, getattr(num,'gt'))
#         elif num.HasField('gte'):
#             if getattr(num, 'lte') > getattr(num, 'gte'):
#                 returnval += '\n  if p.%s < %s or p.%s > %s:\n    return False' %(f.name, getattr(num,'gte'), f.name, getattr(num,'lte'))
#             else:
#                 returnval += '\n  if p.%s > %s and p.%s < %s:\n    return False' %(f.name, getattr(num,'lte'), f.name, getattr(num,'gte'))
#         else:
#             returnval += '\n  if p.%s > %s:\n    return False' %(f.name, getattr(num,'lte'))
#     elif num.HasField('gt'):
#         returnval += '\n  if p.%s <= %s:\n    return False' %(f.name, getattr(num,'gt'))
#     elif num.HasField('gte'):
#         returnval += '\n  if p.%s < %s:\n    return False' %(f.name, getattr(num,'gte'))
#     if getattr(num,'in'):
#         returnval += '\n  if p.%s not in %s:\n    return False' %(f.name, getattr(num,'in'))
#     if getattr(num,'not_in'):
#         returnval += '\n  if p.%s in %s:\n    return False' %(f.name, getattr(num,'not_in'))
#     return returnval
#
# def bool_template(option_value, f):
#     b = option_value.bool
#     returnval = ""
#     if getattr(b,'const') is not None:
#         returnval += '\n  if %s != p.%s:\n    return False' %(getattr(b,'const'), f.name)
#     return returnval
#
# def string_template(option_value, f):
#     s = option_value.string
#     returnval = ""
#     if getattr(s,'const'):
#         returnval += '\n  if \"%s\" != p.%s:\n    return False' %(getattr(s,'const'), f.name)
#     if getattr(s,'in'):
#         returnval += '\n  if p.%s not in %s:\n    return False' %(f.name, getattr(s,'in'))
#     if getattr(s,'not_in'):
#         returnval += '\n  if p.%s in %s:\n    return False' %(f.name, getattr(s,'not_in'))
#     if getattr(s, 'len'):
#         returnval += '\n  if len(p.%s) != %s:\n    return False' %(f.name, getattr(s,'len'))
#     if getattr(s, 'min_len'):
#         returnval += '\n  if len(p.%s) < %s:\n    return False' %(f.name, getattr(s,'min_len'))
#     if getattr(s, 'max_len'):
#         returnval += '\n  if len(p.%s) > %s:\n    return False' %(f.name, getattr(s,'max_len'))
#     if getattr(s, 'len_bytes'):
#         returnval += '\n  if byte_len(p.%s) != %s:\n    return False' %(f.name, getattr(s,'len_bytes'))
#     if getattr(s, 'min_bytes'):
#         returnval += '\n  if byte_len(p.%s) < %s:\n    return False' %(f.name, getattr(s,'min_bytes'))
#     if getattr(s, 'max_bytes'):
#         returnval += '\n  if byte_len(p.%s) > %s:\n    return False' %(f.name, getattr(s,'max_bytes'))
#     if getattr(s, 'pattern'):
#         returnval += '\n  if re.search(r\'%s\', p.%s) is None:\n    return False' %(getattr(s,'pattern'), f.name)
#     if getattr(s, 'prefix'):
#         returnval += '\n  if not p.%s.startswith(\"%s\"):\n    return False' %(f.name, getattr(s,'prefix'))
#     if getattr(s, 'suffix'):
#         returnval += '\n  if not p.%s.endswith(\"%s\"):\n    return False' %(f.name, getattr(s,'suffix'))
#     if getattr(s, 'contains'):
#         returnval += '\n  if \"%s\" not in p.%s:\n    return False' %(getattr(s,'contains'), f.name)
#     if getattr(s, 'email'):
#         returnval += '\n  if not _validateEmail(p.%s):\n    return False' %f.name
#     if getattr(s, 'hostname'):
#         returnval += '\n  if not _validateHostName(p.%s):\n    return False' %f.name
#     if getattr(s, 'address'):
#         returnval += '\n  try:\n    ipaddress.ip_address(unicode(p.%s))\n  except ValueError:\n    if not _validateHostName(p.%s):\n      return False' %(f.name, f.name)
#     if getattr(s, 'ip'):
#         returnval += '\n  try:\n    ipaddress.ip_address(unicode(p.%s))\n  except ValueError:\n    return False' %f.name
#     if getattr(s, 'ipv4'):
#         returnval += '\n  try:\n    ipaddress.IPv4Address(unicode(p.%s))\n  except ValueError:\n    return False' %f.name
#     if getattr(s, 'ipv6'):
#         returnval += '\n  try:\n    ipaddress.IPv6Address(unicode(p.%s))\n  except ValueError:\n    return False' %f.name
#     if getattr(s, 'uri'):
#         returnval += '\n  url = urlparse.urlparse(p.%s)' %f.name
#         returnval += '\n  if not all([url.scheme, url.netloc, url.path]):\n    return False'
#     if getattr(s, 'uri_ref'):
#         returnval += '\n  url = urlparse.urlparse(p.%s)' %f.name
#         returnval += '\n  if not all([url.scheme, url.path]) and url.fragment:\n    return False'
#     if getattr(s, 'uuid'):
#         returnval += '\n  try:\n    uuid.UUID(p.%s)\n  except ValueError:\n    return False' %f.name
#     return returnval

def _has_field(message_pb, property_name):
    # NOTE: As of proto3, HasField() only works for message fields, not for
    #       singular (non-message) fields. First try to use HasField and
    #       if it fails (with a ValueError) we manually consult the fields.
    try:
        return message_pb.HasField(property_name)
    except:
        all_fields = set([field.name for field in message_pb.DESCRIPTOR.fields])
        return property_name in all_fields


def string_template(option_value, f, name):
    str_templ = """{% if s['const'] %}
    if p.{{ name }} != \"{{ s['const'] }}\":
        raise ValidationFailed(\"{{ name }} not equal to {{ s['const'] }}\")
    {% endif %}
    {% if s['in'] %}
    if p.{{ name }} not in {{ s['in'] }}:
        raise ValidationFailed(\"{{ name }} not equal to {{ s['in'] }}\")
    {% endif %}
    {% if s['not_in'] %}
    if p.{{ name }} in {{ s['not_in'] }}:
        raise ValidationFailed(\"{{ name }} in {{ s['not_in'] }}\")
    {% endif %}
    """
    return Template(str_templ).render(s=option_value.string,f=f,name=name)

def message_template(option_value, f, name):
    message_tmpl = """{% if m.message and m.message['required'] %}
    if not _has_field(p, \"{{ name }}\"):
        raise ValidationFailed(\"{{ name }} is required.\")
    {% endif %}
    {% if m.message and m.message['skip'] %}
    # Skipping validation for {{ name }}
    {% else %}
    if _has_field(p, \"{{ name }}\"):
        return generate_validate(p.{{name}})(p.{{name}})    
    {% endif %}
    """
    return Template(message_tmpl).render(m=option_value,f=f, name=name, generate_validate=generate_validate)

def rule_type(field, name = ""):
    if has_validate(field) and field.message_type is None and not field.containing_oneof:
        for option_descriptor, option_value in field.GetOptions().ListFields():
            if option_descriptor.full_name == "validate.rules":
                if str(option_value.string) is not "":
                    return string_template(option_value, field, ".".join([x for x in [name, field.name] if x]))
                elif str(option_value.message) is not "":
                    return message_template(option_value, field, ".".join([x for x in [name, field.name] if x]))
                else:
                    return "raise UnimplementedException()"
    if field.message_type and not field.containing_oneof:
        for option_descriptor, option_value in field.GetOptions().ListFields():
            return message_template(option_value, field, ".".join([x for x in [name, field.name] if x]))
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
    return None, p"""
    return Template(file_tmp).render(rule_type=rule_type, p=proto_message, dir=dir)


class UnimplementedException(Exception):
    pass

class ValidationFailed(Exception):
    pass

def generate_validate(proto_message):
    func = file_template(proto_message)
    exec(func); return validate
