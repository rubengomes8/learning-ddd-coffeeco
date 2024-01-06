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

### Entity vs Value Object
When we can answer _yes_ to all of the following questions, it is a good sign that we can model it as a value object. Also it is better to model something as a value object and then upgrade it to entity later if needed. 
1. Is it possible to treat the object as immutable?
2. Does it measure, quantify, or describe a domain concept?
3. Can it be compared to other objects of the same type just by its values?


## DDD decisions 
1. coffee lover is an entity, because we want them to be defined by their identity. Whenever we are talking to about a coffee lover and adding a CoffeeBux to their loyalty account, there should be no doubt as to which coffee lover we are applying these.
2. store is an entity, because when we reference a store we need to identify which one we are talking about.
3. product answers yes to all the 3 questions, so we will model it as a value object
4. purchase is an entity, it should have an own id. If a customer ever wants a refund on an item, we will need to be able to reference a specific transaction
5. coffee bux is an entity
