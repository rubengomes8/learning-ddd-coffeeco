# learning-ddd-coffeeco
Building a simple monolithic application using domain driven design principles, following the O'Reilly book: https://learning.oreilly.com/library/view/domain-driven-design-with/9781804613450/

## Ubiquitous language
- **Coffee lovers**: What CoffeeCo calls its customers.
- **CoffeeBux**: This is the name of their loyalty program. Coffee lovers earn one CoffeeBux for each drink or accessory they purchase.
- **Tiny, medium, and massive**: The sizes of the drinks are in ascending order. Some drinks are only available in one size, others in all three. Everything on the menu fits into these categories.

## Domains
During the domain modeling session, the following domains were identified:
- Store
- Products
- Loyalty
- Subscription

## Minimum Viable Product (MVP)
- Purchasing a drink or accessory using CoffeeBux
- Purchasing a drink or accessory with a debit/credit card
- Purchasing a drink or accessory with cash
- Earning CoffeeBux on purchases
- Store-specific (but not national) discounts
- We can assume all purchases are in USD for now; in the future, we need to support many currencies though
- Drinks only need to come in one size for now
