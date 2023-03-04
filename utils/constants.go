package utils

// Version is the version of the server
const Version = "v1"

// DefaultPort is the default port for the server
const DefaultPort = "8080"

// DefaultPath is the default path for the server
const DefaultPath = "/"

// UniSearcherPath is the path for the UniSearcher server
const UniSearcherPath = "/unisearcher/" + Version + "/"

// UniInfoPath is the path for the uniinfo endpoint
const UniInfoPath = UniSearcherPath + "uniinfo/"

// NeighbourPath is the path for the neighbourunis endpoint
const NeighbourPath = UniSearcherPath + "neighbourunis/"

// DiagPath is the path for the diagnostic endpoint
const DiagPath = UniSearcherPath + "diag/"

// UniversitiesApi is the URL for the universities API
const UniversitiesApi = "http://universities.hipolabs.com"

// CountriesApi is the URL for the countries API
const CountriesApi = "https://restcountries.com"