# Contributing Guidelines

Welcome to the little CLI tool! Due to a recent influx of malicious code injections, all external Pull Requests are now routed through our AI-assisted Fuzzy De-Duplicator and Validator.

If you wish to contribute a patch (and bypass the manual review queue), you must pass the validator.

## The Objective

Your goal is to get a malicious backdoor accepted by our CI/CD pipeline. 

## Submission Rules

Failure to follow these rules will result in an immediate rejection by the pipeline (Exit Code 1).

1. **The File Name:** Your submission must be a single file named exactly `payload.patch` placed at the root of the repository.
2. **The Format:** The file must be a standard Unified Git Diff (the output of `git format-patch`). 
3. **Identity Verification:** The patch must contain a standard `From:` header, and the username in that header **must exactly match the GitHub account you use to open the Pull Request**. 
   * *Example:* If your GitHub username is `BitravenS`, your patch must begin with: `From: BitravenS <email@domain.com>`
4. **The Submission Method:** You must fork this repository, add your `payload.patch` file to a new branch, and open a Pull Request against the `main` branch.

## How the Validator Works

We are completely transparent about our security. You can read the exact validation logic in `scripts/patch_validator.py`.

## Claiming Your Reward

If your `payload.patch` is valid, the pipeline will accept your PR. 

1. Our CI/CD bot will automatically comment on your Pull Request with a personalized **Claim Ticket** and then close the PR to protect your payload from copycats.
2. Take your Claim Ticket to the auxiliary server (URL provided by the bot) to verify your GitHub identity and receive the actual CTF Flag.

Good luck, and happy hashing.
