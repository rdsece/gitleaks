package gitleaks

const version = "2.0.0"

const defaultGithubURL = "https://api.github.com/"
const defaultThreadNum = 1

// ErrExit used to signal an error during gitleaks execution
const ErrExit = 2

// LeakExit used to signal leaks present in audit
const LeakExit = 1

const defaultConfig = `
title = "sample gitleaks config"

# This is a sample config file for gitleaks. You can configure gitleaks what to search for and what to whitelist.
# The output you are seeing here is the default gitleaks config. If GITLEAKS_CONFIG environment variable
# is set, gitleaks will load configurations from that path. If option --config-path is set, gitleaks will load
# configurations from that path. Gitleaks does not whitelist anything by default.
# References: 
# - https://www.ndss-symposium.org/wp-content/uploads/2019/02/ndss2019_04B-3_Meli_paper.pdf
# - https://github.com/dxa4481/truffleHogRegexes/blob/master/truffleHogRegexes/regexes.json

[[rules]]
description = "AWS Client ID"
regex = '''(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}'''
tags = ["key", "AWS"]

[[rules]]
description = "AWS Secret Key"
regex = '''(?i)aws(.{0,20})?(?-i)['\"][0-9a-zA-Z\/+]{40}['\"]'''
tags = ["key", "AWS"]

[[rules]]
description = "Facebook Secret Key"
regex = '''(?i)(facebook|fb)(.{0,20})?(?-i)['\"][0-9a-f]{32}['\"]'''
tags = ["key", "Facebook"]

[[rules]]
description = "Facebook Client ID"
regex = '''(?i)(facebook|fb)(.{0,20})?['\"][0-9]{13,17}['\"]'''
tags = ["key", "Facebook"]

[[rules]]
description = "Twitter Secret Key"
regex = '''(?i)twitter(.{0,20})?['\"][0-9a-z]{35,44}['\"]'''
tags = ["key", "Twitter"]

[[rules]]
description = "Twitter Client ID"
regex = '''(?i)twitter(.{0,20})?['\"][0-9a-z]{18,25}['\"]'''
tags = ["client", "Twitter"]

[[rules]]
description = "Github Token"
regex = '''(?i)github(.{0,20})?(?-i)['\"][0-9a-zA-Z]{35,40}['\"]'''
tags = ["token", "Github"]

[[rules]]
description = "LinkedIn Client ID"
regex = '''(?i)linkedin(.{0,20})?(?-i)['\"][0-9a-z]{12}['\"]'''
tags = ["client", "Twitter"]

[[rules]]
description = "LinkedIn Secret Key"
regex = '''(?i)linkedin(.{0,20})?['\"][0-9a-z]{16}['\"]'''
tags = ["secret", "Twitter"]

[[rules]]
description = "Square Access Token"
regex = '''(?i)square(.{0,20})?['\"][sq0atp\-[0-9A-Za-z\\-_]]{22}['\"]'''
tags = ["token", "square"]

[[rules]]
description = "Square OAuth Secret"
regex = '''(?i)square(.{0,20})?['\"][sq0csp\-[0-9A-Za-z\\-_]]{43}['\"]'''
tags = ["secret", "square"]

[[rules]]
description = "Twilio API Key"
regex = '''(?i)twilio(.{0,20})?['\"][0-9a-f]{32}['\"]'''
tags = ["secret", "twilio"]

[[rules]]
description = "Stripe API Key"
regex = '''(?i)stripe(.{0,20})?['\"][sk_live[0-9a-z]{24}['\"]'''
tags = ["key", "stripe"]

[[rules]]
description = "Stripe Restricted API Key"
regex = '''(?i)stripe(.{0,20})?['\"][rk_live[0-9a-z]]{24}['\"]'''
tags = ["key", "stripe"]

[[rules]]
description = "MailGun API Key"
regex = '''(?i)(mailgun|mg)(.{0,20})?['\"][0-9a-z]{32}['\"]'''
tags = ["key", "mailgun"]

[[rules]]
description = "MailChimp API Key"
regex = '''(?i)(mailchimp|mc)(.{0,20})?['\"][0-9a-f]{32}-us[0-9]{1,2}['\"]'''
tags = ["key", "mailgun"]

[[rules]]
description = "Heroku API Key"
regex = '''(?i)heroku(.{0,20})?['\"][0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}['\"]'''
tags = ["key", "heroku"]

[[rules]]
description = "GCP API Key"
regex = '''(?i)(google|gcp|youtube|drive|yt)(.{0,20})?['\"][AIza[0-9a-z\\-_]{35}]['\"]'''
tags = ["key", "gcp"]

[[rules]]
description = "GCP OAuth"
regex = '''(?i)(google|gcp|auth)(.{0,20})?['\"][0-9]+-[0-9a-z_]{32}\\.apps\\.googleusercontent\\.com['\"]'''
tags = ["key", "gcp"]

[[rules]]
description = "Slack"
regex = '''xox[baprs]-([0-9a-zA-Z]{10,48})?'''
tags = ["key", "Slack"]

[whitelist]
files = [
  "(.*?)(jpg|gif|doc|pdf|bin)$"
]

#commits = [
#  "whitelisted-commit1",
#  "whitelisted-commit2",
#]
#repos = [
#	"whitelisted-repo"
#]
`
