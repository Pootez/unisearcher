# Assignment 1
Sondre Espeland | *sondesp@ntnu.no* | 03.02.2023

## Specification
### Endpoints
```
/unisearcher/v1/uniinfo/
/unisearcher/v1/neighbourunis/
/unisearcher/v1/diag/
```

The supported request/response pairs are specified in the following.

For the specifications, the following syntax applies:

-   `{:value}` indicates mandatory input parameters specified by the user
-   `{value}` indicates optional input specified by the user, where 'value' can itself contain further optional input. The same notation applies for HTTP parameter specifications (e.g., `{?param}`).

### Retrieve information for a given university

The initial endpoint focuses on return information about a country a particular university/ies (or universities containing a particular string in their name) is/are situated in, such as the official name of the country, spoken languages, and the OpenStreetMap link to the map.

#### Request

```
Method: GET
Path: uniinfo/{:partial_or_complete_university_name}/
```

Note: The name of the university can be partial or complete, and may return a single ("Cambridge") or multiple universities (e.g., "Middle").

Example request: ```uniinfo/norwegian%20university%20of%20science%20and%20technology/```

#### Response

* Content type: `application/json`
* Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):
```
[
  {
      "name": "Norwegian University of Science and Technology", 
      "country": "Norway",
      "isocode": "NO",
      "webpages": ["http://www.ntnu.no/"],
      "languages": {"nno": "Norwegian Nynorsk",
                    "nob": "Norwegian Bokmål",
                    "smi": "Sami"},
      "map": "https://www.openstreetmap.org/relation/2978650"
  },
  ...
]
```

### Retrieve universities with same name components in neighbouring countries

The second endpoint provides an overview of universities in neighbouring countries to a given country that have the same name component (e.g., "Middle") in their institution name. This should **not** include universities from the given country itself.

#### Request

```
Method: GET
Path: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}
```

```{:country_name}``` refers to the English name for the country that is the basis (basis country) of the search of unis with the same name in neighbouring countries.

```{:partial_or_complete_university_name}``` is the partial or complete university name, for which universities with similar name are sought in neighbouring countries

```{?limit={:number}}``` is an optional parameter that limits the number of universities in bordering countries (```number```) that are reported.


Example request: ```neighbourunis/norway/science?limit=5```

#### Response

* Content type: `application/json`
* Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):
```
[
  {
      "name": "Vaasa University of Applied Sciences", 
      "country": "Finland",
      "isocode": "FI",
      "webpages": ["http://www.puv.fi/"],
      "languages": {"fin": "Finnish",
                    "swe": "Swedish"},
      "map": "https://www.openstreetmap.org/relation/54224"
  },
  {
      "name": "Swedish University of Agricultural Sciences", 
      "country": "Sweden",
      "isocode": "SE",
      "webpages": ["http://www.slu.se/"],
      "languages": {"swe":"Swedish"},
      "map": "https://www.openstreetmap.org/relation/52822"
  },
  ...
]
```

### Diagnostics interface

The diagnostics interface indicates the availability of individual services this service depends on. The reporting occurs based on status codes returned by the dependent services, and it further provides information about the uptime of the service.

#### Request

```
Method: GET
Path: diag/
```

### Response

* Content type: `application/json`
* Status code: 200 if everything is OK, appropriate error code otherwise. 

Body:
```
{
   "universitiesapi": "<http status code for universities API>",
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
   "uptime": <time in seconds from the last service restart>
}
```

Note: ```<some value>``` indicates placeholders for values to be populated by the service.
