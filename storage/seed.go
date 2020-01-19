package storage


const serverQuery = `
	INSERT INTO server 
		(name, address, locked)
	VALUES
		(?, ?, ?)
`

func SeedDb(Db Queryer) {
	PreparedExec(
		Db, serverQuery,
		"Vanilluxe", "vanilluxe.th3-z.xyz", 1,
	)

	PreparedExec(
		Db, serverQuery,
		"KF2", "kf2.th3-z.xyz", 0,
	)
}
