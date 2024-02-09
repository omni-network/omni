# Branching Model

Our project adheres to the **Trunk Based Development** approach for managing branches. This methodology emphasizes the use of short-lived development branches, allowing us to integrate "micro-commits" directly into the stable main branch. This practice not only accelerates the review process but also fosters early alignment and increases development velocity.

For an in-depth understanding of Trunk Based Development, we recommend this resource: [TrunkBasedDevelopment.com](https://trunkbaseddevelopment.com/).

Key principles of our branching model include:

- **Short-Lived Branches**: Development branches are kept brief in lifespan to facilitate quick integrations and reviews.
- **Direct Commits to Main**: Micro-commits are merged directly into the main branch to ensure that it remains stable and up-to-date.
- **Avoid Long-Lived Feature Branches**: To prevent the complexities of merge conflicts and integration challenges, large feature developments are divided into smaller, incremental commits.

## Branch Naming

Branch names should be descriptive and follow a consistent naming convention. We recommend using the following format:

```
<author>/<description>
```
