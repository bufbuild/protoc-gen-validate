package tpl

const repTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}{{ $typ := inType $f nil }}

	{{ if $r.GetMinItems }}
		{{ if eq $r.GetMinItems $r.GetMaxItems }}
			if ({{ accessor . }}.size() != {{ $r.GetMinItems }}) {
				{{ err . "value must contain exactly " $r.GetMinItems " item(s)" }}
			}
		{{ else if $r.MaxItems }}
			if ({{ accessor . }}.size() < {{ $r.GetMinItems }} || {{ accessor . }}.size() > {{ $r.GetMaxItems }}) {
			 	{{ err . "value must contain between " $r.GetMinItems " and " $r.GetMaxItems " items, inclusive" }}
			}
		{{ else }}
			if ({{ accessor . }}.size() < {{ $r.GetMinItems }}) {
				{{ err . "value must contain at least " $r.GetMinItems " item(s)" }}
			}
		{{ end }}
	{{ else if $r.MaxItems }}
		if ({{ accessor . }}.size() > {{ $r.GetMaxItems }}) {
			{{ err . "value must contain no more than " $r.GetMaxItems " item(s)" }}
		}
	{{ end }}

	{{ if $r.GetUnique }}
	// Implement comparison for wrapped reference types
	struct cmp {
		bool operator() (const std::reference_wrapper<{{ $typ }}> lhs, const std::reference_wrapper<{{ $typ }}> rhs) const {
			return lhs.get() < rhs.get();
		}
	};

	// Save a set of references to avoid copying overhead
	std::set<std::reference_wrapper<{{ $typ }}>,cmp> {{ lookup $f "Unique" }};
	{{ end }}

	{{ if or $r.GetUnique (ne (.Elem "" "").Typ "none") }}
		for (int i = 0; i < {{ accessor . }}.size(); i++) {
			const {{ $typ }}& item = {{ accessor . }}.Get(i);

			{{ if $r.GetUnique }}
				auto p = {{ lookup $f "Unique" }}.emplace(const_cast<{{ $typ }}&>(item));
				if (p.second == false) {
					{{ errIdx . "idx" "repeated value must contain unique items" }}
				}
			{{ end }}

			{{ render (.Elem "item" "i") }}
		}
	{{ end }}
`
