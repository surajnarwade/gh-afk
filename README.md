# gh-afk

`gh-afk` is a GitHub CLI extension that helps you manage your GitHub user status directly from the terminal. 
Set yourself as AFK (Away From Keyboard) to prevent automatic pull request review assignments when you're in meetings, on vacation, or otherwise unavailable.

## Why this Extension?

In some organizations, pull request reviews are automatically requested and assigned to team members. Sometimes you're in a meeting, out sick, or on vacation and forget to update your GitHub status. When your status is set to busy, GitHub won't assign reviews to you, which helps manage your workload.

If you have the GitHub mobile app, you can update your status on the go. But why not do it directly from the terminal? Since there's no REST API but a GraphQL API is available, this extension makes it easy.

This is a fun Friday project to help solve this issue. Feel free to star it⭐, fork it, use it, and if you see any issues, send a PR!

## Installation

It will only work if you use gh CLI, you can install it with gh CLI

```
gh extensions install surajnarwade/gh-afk
```

## Usage

* Set the AFK status

```
gh afk on
```

By default, it will use:
* emoji: :no_entry:
* message: AFK

you can override this setting using flags:

```
gh afk on -msg "your-message" -emoji ":your_emoji"
```

* Remove the AFK status

```
gh afk off
```

* View AFK status

```
gh afk view
```

| Note: Currently this will set Expiration to `Never` so you will have to manually turn the status on or off.

## Troubleshooting

If you see following or similar error:

```
{
  "errors": [
    {
      "type": "INSUFFICIENT_SCOPES",
      "locations": [
        {
          "line": 1,
          "column": 10
        }
      ],
      "message": "Your token has not been granted the required scopes to execute this query. The 'changeUserStatus' field requires one of the following scopes: ['user'], but your token has only been granted the: ['admin:public_key', 'gist', 'read:org', 'repo'] scopes. Please modify your token's scopes at: https://github.com/settings/tokens."
    }
  ]
}
gh: Your token has not been granted the required scopes to execute this query. The 'changeUserStatus' field requires one of the following scopes: ['user'], but your token has only been granted the: ['admin:public_key', 'gist', 'read:org', 'repo'] scopes. Please modify your token's scopes at: https://github.com/settings/tokens.
```

quick resolution would be to run the following command,

```
gh auth refresh -s user
```

