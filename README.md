# OpenRSPP

An Open-Source Platform for Finding and Comparing Research Service Providers

## Introduction

OpenRSPP Marketplace is an open-source platform for finding and comparing research service providers (RSPPs). The platform provides a centralized location for researchers and scientific organizations to discover and compare RSPPs, and find the right service provider for their needs.

## Key Features

* Centralized platform for finding and comparing RSPPs
* Detailed information on RSPPs, including services offered and pricing
* arch and filter capabilities to easily find the right RSPP for your needs
* User ratings and reviews to help you make informed decisions
* Open-source and community-driven, with contributions from RSPPs and other organizations

## Service Provider Management
### Get Service Provider Information
`GET /serviceproviders/:id`

Retrieve information about a specific service provider, including services offered and pricing.

Request

```
curl -X GET https://api.openrspp.io/serviceproviders/:id
```

Response

```
HTTP/1.1 200 OK
{
    "id": 1,
    "name": "RSPP A",
    "services": [
        {
            "id": 1,
            "name": "Clinical Trial Design",
            "price": 5000
        },
        {
            "id": 2,
            "name": "Data Management",
            "price": 4000
        }
    ],
    "rating": 4.5,
    "review_count": 10
}
```

### Create a Service Provider

`POST /serviceproviders`

Create a new service provider, including services and pricing.

Request

```
curl -X POST https://api.openrspp.io/serviceproviders
```
```
{
    "name": "RSPP B",
    "services": [
        {
            "name": "Clinical Trial Design",
            "price": 5000
        },
        {
            "name": "Data Management",
            "price": 4000
        }
    ]
}
```
Response
```
HTTP/1.1 201 Created
{
    "id": 2,
    "name": "RSPP B",
    "services": [
        {
            "id": 3,
            "name": "Clinical Trial Design",
            "price": 5000
        },
        {
            "id": 4,
            "name": "Data Management",
            "price": 4000
        }
    ],
    "rating": 0,
    "review_count": 0
}
```

## Service Management
### Update Service Information
`PUT /services/:id`

Update the information for a specific service, including price.

Request

```
curl -X PUT https://api.openrspp.io/services/:id
```

```
{
    "price": 6000
}
```

Response

```
HTTP/1.1 200 OK
{
    "id": 1,
    "name": "Clinical Trial Design",
    "price": 6000
}
```

Add Service
POST /services

Add a new service to a service provider.

Request
```
curl -X POST https://api.openrspp.io/services
```

```
{
    "name": "Regulatory Support",
    "price": 7000
}
```
Response
```
HTTP/1.1 201 Created
{
    "id": 5,
    "name": "Regulatory Support",
    "price": 7000
}
```

## Search and Filter
### Search Service Providers

`GET /serviceproviders`

Search and filter service providers based on name, services offered, price, rating, and review count.

Request
```
curl -X GET https://api.openrspp.io/serviceproviders?name=RSPP&services=Clinical%20Trial%20Design&min_price=3000&max_price=6000&min_rating=4&max_rating=5&min_review_count=5
```
Response
```
HTTP/1.1 200 OK
[    {        "id": 1,        "name": "RSPP A",        "services": [            {                "id": 1,                "name": "Clinical Trial Design",                "price": 5000            }        ],
        "rating": 4.8,
        "review_count": 8
    },
    {
        "id": 2,
        "name": "RSPP B",
        "services": [
            {
                "id": 3,
                "name": "Clinical Trial Design",
                "price": 5000
            }
        ],
        "rating": 4.5,
        "review_count": 10
    }
]
```
## Review Management
### Add Review
POST /reviews

Add a review for a service provider.

Request
```
curl -X POST https://api.openrspp.io/reviews
```
```
{
    "service_provider_id": 1,
    "rating": 4,
    "text": "Great service and excellent support!"
}
```
Response
```
HTTP/1.1 201 Created
{
    "id": 1,
    "service_provider_id": 1,
    "rating": 4,
    "text": "Great service and excellent support!",
    "created_at": "2022-12-01T12:00:00Z"
}
```

### Get Review Count
`GET /serviceproviders/:id/review_count`

Retrieve the number of reviews for a specific service provider.

Request
```
curl -X GET https://api.openrspp.io/serviceproviders/1/review_count
```

Response
```
HTTP/1.1 200 OK
{
    "review_count": 8
}
```


## Getting Started

To get started with OpenRSPP Marketplace, simply download the software and set it up on a server or cloud-based platform of your choice. You can then host your own network of RSPPs and connect with researchers and organizations looking for services.

If you're an RSPP, you can join a network hosted by a company using OpenRSPP Marketplace and start showcasing your services to researchers and organizations. OpenRSPP Marketplace provides a platform for attracting new clients and growing your business.

## Contributions

OpenRSPP Marketplace is an open-source and community-driven project, and we welcome contributions from RSPPs and other organizations. Whether you're a developer, researcher, or service provider, you can help advance the platform and make it better for everyone.

To contribute, simply fork the repository and submit a pull request with your changes. You can also join the community and participate in discussions and forums.

## Support

For support and assistance with OpenRSPP Marketplace, please consult the documentation and join the community. You can also reach out to the project maintainers for additional help.

## License

OpenRSPP Marketplace is open-source software released under the MIT License.
