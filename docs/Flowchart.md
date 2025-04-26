flowchart TD
A[Checkout] --> B{Diskon}
B -->|Yes| C[Calcalating Diskon]
C --> D[Total Payment]
B -->|No| D