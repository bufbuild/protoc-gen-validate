import re
from validate_email import validate_email
import ipaddress
try:
    import urlparse
except ImportError:
    import urllib.parse as urlparse
import uuid
import struct

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

def num_template(option_value, f, num):
    returnval = ""
    if str(option_value.float) is not "": # Need to convert double to float
        if num.HasField('const'):
            returnval += '\n  if struct.unpack(\"f\", struct.pack(\"f\", %s))[0] != p.%s:\n    return False' %(getattr(num,'const'), f.name)
    else:
        if num.HasField('const'):
            returnval += '\n  if %s != p.%s:\n    return False' %(getattr(num,'const'), f.name)
    if num.HasField('lt'):
        if num.HasField('gt'):
            if getattr(num, 'lt') > getattr(num, 'gt'):
                returnval += '\n  if p.%s <= %s or p.%s >= %s:\n    return False' %(f.name, getattr(num,'gt'), f.name, getattr(num,'lt'))
            else:
                returnval += '\n  if p.%s >= %s and p.%s <= %s:\n    return False' %(f.name, getattr(num,'lt'), f.name, getattr(num,'gt'))
        elif num.HasField('gte'):
            if getattr(num, 'lt') > getattr(num, 'gte'):
                returnval += '\n  if p.%s < %s or p.%s >= %s:\n    return False' %(f.name, getattr(num,'gte'), f.name, getattr(num,'lt'))
            else:
                returnval += '\n  if p.%s >= %s and p.%s < %s:\n    return False' %(f.name, getattr(num,'lt'), f.name, getattr(num,'gte'))
        else:
            returnval += '\n  if p.%s >= struct.unpack(\"f\", struct.pack(\"f\", %s))[0]:\n    return False' %(f.name, getattr(num,'lt'))
    elif num.HasField('lte'):
        if num.HasField('gt'):
            if getattr(num, 'lte') > getattr(num, 'gt'):
                returnval += '\n  if p.%s <= %s or p.%s > %s:\n    return False' %(f.name, getattr(num,'gt'), f.name, getattr(num,'lte'))
            else:
                returnval += '\n  if p.%s > %s and p.%s <= %s:\n    return False' %(f.name, getattr(num,'lte'), f.name, getattr(num,'gt'))
        elif num.HasField('gte'):
            if getattr(num, 'lte') > getattr(num, 'gte'):
                returnval += '\n  if p.%s < %s or p.%s > %s:\n    return False' %(f.name, getattr(num,'gte'), f.name, getattr(num,'lte'))
            else:
                returnval += '\n  if p.%s > %s and p.%s < %s:\n    return False' %(f.name, getattr(num,'lte'), f.name, getattr(num,'gte'))
        else:
            returnval += '\n  if p.%s > %s:\n    return False' %(f.name, getattr(num,'lte'))
    elif num.HasField('gt'):
        returnval += '\n  if p.%s <= %s:\n    return False' %(f.name, getattr(num,'gt'))
    elif num.HasField('gte'):
        returnval += '\n  if p.%s < %s:\n    return False' %(f.name, getattr(num,'gte'))
    if getattr(num,'in'):
        returnval += '\n  if p.%s not in %s:\n    return False' %(f.name, getattr(num,'in'))
    if getattr(num,'not_in'):
        returnval += '\n  if p.%s in %s:\n    return False' %(f.name, getattr(num,'not_in'))
    return returnval

def bool_template(option_value, f):
    b = option_value.bool
    returnval = ""
    if getattr(b,'const') is not None:
        returnval += '\n  if %s != p.%s:\n    return False' %(getattr(b,'const'), f.name)
    return returnval

def string_template(option_value, f):
    s = option_value.string
    returnval = ""
    if getattr(s,'const'):
        returnval += '\n  if \"%s\" != p.%s:\n    return False' %(getattr(s,'const'), f.name)
    if getattr(s,'in'):
        returnval += '\n  if p.%s not in %s:\n    return False' %(f.name, getattr(s,'in'))
    if getattr(s,'not_in'):
        returnval += '\n  if p.%s in %s:\n    return False' %(f.name, getattr(s,'not_in'))
    if getattr(s, 'len'):
        returnval += '\n  if len(p.%s) != %s:\n    return False' %(f.name, getattr(s,'len'))
    if getattr(s, 'min_len'):
        returnval += '\n  if len(p.%s) < %s:\n    return False' %(f.name, getattr(s,'min_len'))
    if getattr(s, 'max_len'):
        returnval += '\n  if len(p.%s) > %s:\n    return False' %(f.name, getattr(s,'max_len'))
    if getattr(s, 'len_bytes'):
        returnval += '\n  if byte_len(p.%s) != %s:\n    return False' %(f.name, getattr(s,'len_bytes'))
    if getattr(s, 'min_bytes'):
        returnval += '\n  if byte_len(p.%s) < %s:\n    return False' %(f.name, getattr(s,'min_bytes'))
    if getattr(s, 'max_bytes'):
        returnval += '\n  if byte_len(p.%s) > %s:\n    return False' %(f.name, getattr(s,'max_bytes'))
    if getattr(s, 'pattern'):
        returnval += '\n  if re.search(r\'%s\', p.%s) is None:\n    return False' %(getattr(s,'pattern'), f.name)
    if getattr(s, 'prefix'):
        returnval += '\n  if not p.%s.startswith(\"%s\"):\n    return False' %(f.name, getattr(s,'prefix'))
    if getattr(s, 'suffix'):
        returnval += '\n  if not p.%s.endswith(\"%s\"):\n    return False' %(f.name, getattr(s,'suffix'))
    if getattr(s, 'contains'):
        returnval += '\n  if \"%s\" not in p.%s:\n    return False' %(getattr(s,'contains'), f.name)
    if getattr(s, 'email'):
        returnval += '\n  if not _validateEmail(p.%s):\n    return False' %f.name
    if getattr(s, 'hostname'):
        returnval += '\n  if not _validateHostName(p.%s):\n    return False' %f.name
    if getattr(s, 'address'):
        returnval += '\n  try:\n    ipaddress.ip_address(unicode(p.%s))\n  except ValueError:\n    if not _validateHostName(p.%s):\n      return False' %(f.name, f.name)
    if getattr(s, 'ip'):
        returnval += '\n  try:\n    ipaddress.ip_address(unicode(p.%s))\n  except ValueError:\n    return False' %f.name
    if getattr(s, 'ipv4'):
        returnval += '\n  try:\n    ipaddress.IPv4Address(unicode(p.%s))\n  except ValueError:\n    return False' %f.name
    if getattr(s, 'ipv6'):
        returnval += '\n  try:\n    ipaddress.IPv6Address(unicode(p.%s))\n  except ValueError:\n    return False' %f.name
    if getattr(s, 'uri'):
        returnval += '\n  url = urlparse.urlparse(p.%s)' %f.name
        returnval += '\n  if not all([url.scheme, url.netloc, url.path]):\n    return False'
    if getattr(s, 'uri_ref'):
        returnval += '\n  url = urlparse.urlparse(p.%s)' %f.name
        returnval += '\n  if not all([url.scheme, url.path]) and url.fragment:\n    return False'
    if getattr(s, 'uuid'):
        returnval += '\n  try:\n    uuid.UUID(p.%s)\n  except ValueError:\n    return False' %f.name
    return returnval

class UnimplementedException(Exception):
    pass

def generate_validate(proto_message):
    file = ""
    func = file + '\ndef validate(p): '
    disabled = False
    for options, values in proto_message.DESCRIPTOR.GetOptions().ListFields():
        if options.full_name == "validate.disabled" and values == True:
            disabled = True
    if not disabled:
        for field in proto_message.DESCRIPTOR.fields:
            if has_validate(field) and field.message_type is None and not field.containing_oneof:
                for option_descriptor, option_value in field.GetOptions().ListFields():
                    if option_descriptor.full_name == "validate.rules":
                        if str(option_value.string) is not "":
                            func += string_template(option_value, field)
                        elif str(option_value.bool) is not "":
                            func += bool_template(option_value, field)
                        elif str(option_value.float) is not "":
                            func += num_template(option_value, field, option_value.float)
                        elif str(option_value.double) is not "":
                            func += num_template(option_value, field, option_value.double)
                        elif str(option_value.int32) is not "":
                            func += num_template(option_value, field, option_value.int32)
                        elif str(option_value.int64) is not "":
                            func += num_template(option_value, field, option_value.int64)
                        elif str(option_value.uint32) is not "":
                            func += num_template(option_value, field, option_value.uint32)
                        elif str(option_value.uint64) is not "":
                            func += num_template(option_value, field, option_value.uint64)
                        elif str(option_value.sint32) is not "":
                            func += num_template(option_value, field, option_value.sint32)
                        elif str(option_value.sint64) is not "":
                            func += num_template(option_value, field, option_value.sint64)
                        elif str(option_value.fixed32) is not "":
                            func += num_template(option_value, field, option_value.fixed32)
                        elif str(option_value.fixed64) is not "":
                            func += num_template(option_value, field, option_value.fixed64)
                        elif str(option_value.sfixed32) is not "":
                            func += num_template(option_value, field, option_value.sfixed32)
                        elif str(option_value.sfixed64) is not "":
                            func += num_template(option_value, field, option_value.sfixed64)
                        else:
                            func += "\n  raise UnimplementedException()"
            elif field.message_type or field.containing_oneof:
                func += "\n  raise UnimplementedException()"
    func += "\n  return True"
    exec(func); return validate
