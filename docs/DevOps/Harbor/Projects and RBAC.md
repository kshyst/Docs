# Projects and RBAC

Harbor uses **Projects** as the primary organizational unit to manage repositories and access control.

## Project Types

1. **Public Projects**: Anyone can pull images from a public project, even without logging in. Only members of the project can push images.
2. **Private Projects**: Only authenticated users who are members of the project can pull or push images.

## Managing Members

You can add users or groups (if LDAP/OIDC is integrated) to a project and assign them specific roles.

### User Roles

- **Project Admin**: Has full control over the project, including managing members, repositories, and project settings.
- **Maintainer**: Can manage repositories, pull/push images, and manage Helm charts, but cannot manage members or project settings.
- **Developer**: Can pull and push images and view project information.
- **Guest**: Can pull images and view project information but cannot push images.
- **Limited Guest**: Can pull images but has limited view of project information (cannot see logs or members).

## Access Control Lists (ACL)

RBAC in Harbor ensures that permissions are strictly enforced. For example:
- A user with the **Developer** role in Project A cannot push images to Project B unless they are also a member of Project B with push permissions.
- **System Admins** have full access to all projects and system-wide settings.

## Robot Accounts

For automated processes like CI/CD pipelines, Harbor provides **Robot Accounts**. These accounts are not associated with a real person and can be granted specific permissions (pull, push, or both) to one or more projects. They use tokens for authentication instead of passwords.
