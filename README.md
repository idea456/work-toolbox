## Usage/Examples

To ping about pull request comments, approvals, rejections etc:

```bash
work-toolbox notify --type pull-request
```

To ping about build progresses (CircleCI, Vercel):

```bash
work-toolbox notify --type builds
```

To ping about Google calendar meetings and standups:

```bash
work-toolbox notify --type calendar
```

## TODO

-   Add command to automate Git workflow:
    -   Run tests, git add, commit and create new pull request and redirect to Github PR page: `work-toolbox commit -m "Fixed issue with tooltip" --pr --test`
-   Add command to add reminders:
    -   `work-toolbox remind --by today -m "Fix that freaking tooltip error once block user card is moved to reviews!`
-   Add command to start calls
