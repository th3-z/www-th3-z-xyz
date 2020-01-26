package storage

import (
	"database/sql"
	"github.com/th3-z/malgo"
)

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

const projectQuery = `
	INSERT INTO project
		(name, project_url, repo_url, description, status)
	VALUES
		(?, ?, ?, ?, ?)
`

func SeedDb(db *sql.DB) {
	PreparedExec(
		db, serverQuery,
		"Vanilluxe", "Minecraft","vanilluxe.th3-z.xyz","https://vanilluxe.th3-z.xyz", 1, 4,
	)

	PreparedExec(
		db, serverQuery,
		"KF2-MA Dev Server", "Killing Floor 2", "kf2.th3-z.xyz", "https://kf2.th3-z.xyz", 1, 6,
	)

	PreparedExec(
		db, infrastructureQuery,
		"beta.th3-z.xyz", "136.244.96.98", "Debian 10",
	)

	PreparedExec(
		db, infrastructureQuery,
		"atlus.th3-z.xyz", "45.32.187.80", "Debian 10",
	)

	PreparedExec(
		db, infrastructureQuery,
		"saturn.th3-z.xyz", "212.159.110.214", "KDE Neon",
	)

	PreparedExec(
		db, projectQuery,
		"Killing Floor 2 Magicked Admin", "https://kf2-ma.th3-z.xyz/", "https://github.com/th3-z/kf2-magicked-admin", "Scripted management, statistics, and bot for Killing Floor 2", "Released",
	)

	PreparedExec(
		db, projectQuery,
		"Malgo", "", "https://github.com/th3-z/malgo", "MyAnimeList SQL migration utility and Go library ", "Released",
	)

	PreparedExec(
		db, projectQuery,
		"Foobar2000 Mini-Player", "", "https://github.com/th3-z/fb2k-mini-player", "A tiny Foobar 2000 controller that can sit above fullscreen applications", "Released",
	)

	PreparedExec(
		db, projectQuery,
		"Zerve", "", "https://github.com/th3-z/zerve", "KDE - File sharing over HTTP from your desktop or cli", "Released",
	)

	PreparedExec(
		db, projectQuery,
		"QRss", "", "https://github.com/th3-z/qrss", "A Qt5 based RSS reader", "Development",
	)

	PreparedExec(
		db, projectQuery,
		"Public website", "", "https://github.com/th3-z/beta-th3-z-xyz", "You're looking at it", "Development",
	)

	malgo.MigrateFile(db, "storage/th3-z-anime-list.xml")
}
