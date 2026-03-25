# Query Language

Query language is built from the ground up to be a powerful and flexible query language that can be used to query data
from various sources. It is written in Go and designed to integrate seamlessly with the QL-database, providing a
unified interface for both ingested and real-time data access.

## Data Servings

Data is served in two ways: ingested data and real-time data. Ingested data is data that has been collected, structured
and stored in the QL-database, while real-time data is data that is being collected and processed in real-time.
Choosing the right serving strategy depends on the use case - ingested data favors consistency and performance, while
real-time data favors freshness and immediacy.

### Ingested Data (Recommended)

Ingested data is data that has been collected, structured and stored in the QL-database. This is the recommended
approach for querying data, as it allows for more efficient and reliable access to pre-processed information.

Because the data has already been validated and indexed at ingestion time, queries against ingested data benefit from
faster execution, lower resource consumption, and predictable response times. Use this serving when working with
historical records, aggregated metrics, or any scenario where data completeness is more important than immediacy.

### Real-time Data

Real-time data is data that is being collected and processed in real-time. This approach is suited for querying data
from sources as it is actively being gathered, without prior storage in the QL-database.

Since the data bypasses the ingestion pipeline, queries may encounter incomplete or unstructured results. This serving
is best used for live monitoring, streaming analytics, or situations where the most up-to-date snapshot of the data is
critical, even at the cost of additional processing overhead.
