package storage

import (
	"database/sql"

	"github.com/th3-z/malgo"
)

const serverQuery = `
	INSERT INTO server 
		(name, service, address, image, web_url, locked, max_players)
	VALUES
		(?, ?, ?, ?, ?, ?, ?)
`

const infrastructureQuery = `
	INSERT INTO infrastructure
		(hostname, address, os)
	VALUES
		(?, ?, ?)
`

const softwareQuery = `
	INSERT INTO software
		(name, project_url, repo_url, description, status)
	VALUES
		(?, ?, ?, ?, ?)
`

func SeedDb(db *sql.DB) {
	PreparedExec(
		db, serverQuery,
		"Vanilluxe MC", "Minecraft", "vanilluxe.th3-z.xyz", "images/servers/mc.png", "https://vanilluxe.th3-z.xyz", 1, 6,
	)
	PreparedExec(
		db, serverQuery,
		"Magicked Admin Dev Server (NL)", "Killing Floor 2", "kf2.th3-z.xyz", "images/servers/kf2.png", "https://kf2.th3-z.xyz", 0, 6,
	)
	PreparedExec(
		db, serverQuery,
		"Magicked Admin Dev Server (UK)", "Killing Floor 2", "saturn.th3-z.xyz", "images/servers/kf2.png", "https://saturn.th3-z.xyz", 0, 6,
	)
	PreparedExec(
		db, serverQuery,
		"Git", "Git", "git.th3-z.xyz", "images/servers/git.png", "https://git.th3-z.xyz", 0, 0,
	)

	PreparedExec(
		db, infrastructureQuery,
		"atlus.th3-z.xyz", "136.244.96.98", "Debian 10",
	)
	PreparedExec(
		db, infrastructureQuery,
		"saturn.th3-z.xyz", "212.159.110.214", "Ubuntu 20.04",
	)

	PreparedExec(
		db, softwareQuery,
		"Killing Floor 2 Magicked Admin", "https://kf2-ma.th3-z.xyz/", "https://github.com/th3-z/kf2-magicked-admin", "Scripted management, statistics, and bot for Killing Floor 2", "Released",
	)
	PreparedExec(
		db, softwareQuery,
		"Gloom", "", "https://github.com/th3-z/gloom-client", "File upload utility and server", "Development",
	)
	PreparedExec(
		db, softwareQuery,
		"Malgo", "", "https://github.com/th3-z/malgo", "MyAnimeList SQL migration utility and Go library ", "Released",
	)
	PreparedExec(
		db, softwareQuery,
		"Foobar2000 Mini-Player", "", "https://github.com/th3-z/fb2k-mini-player", "A tiny Foobar 2000 controller that can sit above fullscreen applications", "Released",
	)
	PreparedExec(
		db, softwareQuery,
		"Zerve", "", "https://github.com/th3-z/zerve", "KDE - File sharing over HTTP from your desktop or cli", "Released",
	)
	PreparedExec(
		db, softwareQuery,
		"QRss", "", "https://github.com/th3-z/qrss", "A Qt5 based RSS reader", "Development",
	)
	PreparedExec(
		db, softwareQuery,
		"Public website", "", "https://github.com/th3-z/www-th3-z-xyz", "This website you're currently viewing", "Development",
	)

	malgo.MigrateFile(db, "storage/th3-z-anime-list.xml")
}
