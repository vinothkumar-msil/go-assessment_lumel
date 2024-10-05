# go-assessment_lumel
assessment

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Go**: Version **1.18** or later
- **PostgreSQL**: Version **13** or later
- **Any other dependencies**: Ensure you have the necessary Go modules installed.

## Getting Started

Follow these steps to set up and run the project locally:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/go-backend-assessment.git
   cd go-backend-assessment



```markdown
## API Endpoints

| Route      | Method | Body                             | Sample Response                                    | Description                                |
|------------|--------|----------------------------------|---------------------------------------------------|--------------------------------------------|
| `/refresh` | GET   | None                             | `{"message": "Data refresh initiated successfully."}` | Triggers a data refresh from the CSV file. |
| `/revenue` | GET    | `start` & `end` (query params) | `{"total_revenue": 1000.00}`                      | Calculates total revenue in a date range.  |
