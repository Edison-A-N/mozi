# Git Commit & Push

## Task
Create a concise, single-line commit message in English based on the staged changes, following the project's commit log format conventions.

## What the AI will do:

1. ✅ Check `git status` to identify staged files
2. ✅ Read staged files to understand the changes
3. ✅ Review `git log --oneline -10` to match existing commit message format
4. ✅ Generate an appropriate commit message in English
5. ✅ Execute `git commit -m "message"`
6. ✅ If pre-commit hooks fail:
   - **For auto-formatting issues**: Run `git diff` to check the auto-formatted changes, then precisely `git add` those files and retry commit
   - **For other issues** (tests, linting errors, etc.): Stop immediately and provide concise guidance on what to do next
7. ✅ Push changes with `git push` (if requested)

## Important Notes
- **Handle formatting auto-fixes**: If pre-commit hooks auto-format code (e.g., prettier, gofmt, black), automatically stage the formatted files and retry commit
- **Stop for real errors**: If hooks fail due to tests, build errors, or other non-formatting issues, stop and clearly explain what needs to be fixed
- **Never** stage unrelated files - only add files that were auto-formatted by the hooks during the commit attempt
