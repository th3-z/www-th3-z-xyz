package storage


const serverQuery = `
	INSERT INTO server 
		(name, service, address, web_url, locked, max_players)
	VALUES
		(?, ?, ?, ?, ?, ?)
`

const infrastructureQuery = `
	INSERT INTO infrastructure
		(hostname, address, os)
	VALUES
		(?, ?, ?)
`

func SeedDb(Db Queryer) {
	PreparedExec(
		Db, serverQuery,
		"Vanilluxe", "Minecraft","vanilluxe.th3-z.xyz","https://vanilluxe.th3-z.xyz", 1, 4,
	)

	PreparedExec(
		Db, serverQuery,
		"KF2-MA Dev Server", "Killing Floor 2", "kf2.th3-z.xyz", "https://kf2.th3-z.xyz", 1, 6,
	)

	PreparedExec(
		Db, infrastructureQuery,
		"beta.th3-z.xyz", "136.244.96.98", "Debian 10",
	)

	PreparedExec(
		Db, infrastructureQuery,
		"atlus.th3-z.xyz", "45.32.187.80", "Debian 10",
	)

	PreparedExec(
		Db, infrastructureQuery,
		"saturn.th3-z.xyz", "212.159.110.214", "KDE Neon",
	)

}
