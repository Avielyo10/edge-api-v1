# Edge API V1

This project uses `docker` to run the clean env.

## Prepare for development

This will install all the dependencies and test your current state.

```bash
make
```

## Clean dev environment

This will remove postgresql and redis containers, then create new ones and start edge-api.

```bash
make clean_env
```

## Why do we need that?

1. Remove DB logic from business logic without getting provider lock.
2. 'Contract based' development (write the spec, then generate code from that spec instead of doing the opposite).
3. CQRS.
4. Domain Driven Design.

## What's is missing?

* [ ] Add missing tests
* [ ] Change relative paths on dockerfiles
* [ ] Design the update service similar to the edge service
* [ ] Create the `Devices` part under the edge service
* [ ] Add the update logic for the edge service (both images & devices)
