package python

const hostTpl = `
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

	return True, ""
`

const emailTpl = `
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
`
