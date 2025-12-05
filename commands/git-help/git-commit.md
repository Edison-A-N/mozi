# Git Commit & Push

## Task
Create a concise, single-line commit message in English based on the staged changes, following the project's commit log format conventions.

## What the AI will do:

1. ✅ Check `git status` to identify staged files
2. ✅ Read staged files to understand the changes
3. ✅ Review `git log --oneline -10` to match existing commit message format
4. ✅ Generate an appropriate commit message in English
5. ✅ Execute `git commit -m "message"`
6. ✅ Push changes with `git push` (if requested)

## Important Notes
- If pre-commit hooks fail, **stop immediately** and notify the user to fix the issues
- **Never** stage or add files automatically - only commit what is already staged
