# GO project boirleplate /Work In Progress/

It is boilerplate that can be used for implementing rest-apis in GO-lang.

## Project layout

https://github.com/golang-standards/project-layout

## Testing

### Unit tests

Unit tests are in standalone directory for few reasons:

- it is cleaner,
- we should test only public methods, so it doesn't need to have access to private one,

## End to end

TODO

## Internal

Components of internal directory.

### DAL

Data access layer, in modules should only be interfaces for repositories, here is the place where access to data can be implemented.

### Modules

Only domain logic.

### Api-v1

All related to handle rest api:

- controllers
- formatters between domain and api models,
- routing

## Pkg

All public resources, like API models.