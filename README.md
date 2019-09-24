# vigilant-barnacle [WIP]

```bash
  // trips-log data file
	pf := "./assets/all_data.csv"
  // Load file
	res := p.ScanDoc(pf)
  
  // Trip states you're matching
	a := []string{"16", "20"}

  // Output
	fmt.Println(p.Filter(res, a))

```

