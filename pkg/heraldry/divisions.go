package heraldry

var (
	Bend               = Division{"bend", "Per bend ", Tincture{}}
	BendSinister       = Division{"bendsinister", "Per bend sinister ", Tincture{}}
	Fess               = Division{"fess", "Per fess ", Tincture{}}
	Pale               = Division{"pale", "Per pale ", Tincture{}}
	Plain              = Division{"plain", "", Tincture{}}
	Quarterly          = Division{"quarterly", "Quarterly ", Tincture{}}
	Saltire            = Division{"saltire", "Per saltire ", Tincture{}}
	Chevron            = Division{"chevron", "Per chevron ", Tincture{}}
	AvailableDivisions = [8]Division{Bend, BendSinister, Fess, Pale, Plain, Quarterly, Saltire, Chevron}
)
