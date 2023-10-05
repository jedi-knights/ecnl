# Angular Setup

This document describes how to setup Angular for local development.

When working with Angular it is beneficial to have the [Node Version Manager](node.md) installed to manage multiple versions of Node.js.
So if you haven't setup NVM yet head to that document first before you start messing around with Angular setup (you can thank me later).

## Install Angular CLI

The Angular CLI is a command-line interface tool that you use to initialize, develop, scaffold, and maintain Angular applications.

As the Angular toolset is rapidly evolving it is recommended to take a quick look at the [Angular CLI](https://cli.angular.io/) 
documentation to see if there are any additional steps required to setup Angular.

You will also need to be aware of the versions of Node.js and Angular CLI that are compatible with each other.  You can find this
on the [version compatibility](https://angular.io/guide/versions) page.  As of this writing the version compatibility doc indicates
that Angular 16.1.x || 16.2.x is compatible with Node.js ^16.14.0 || ^18.10.0.  Seeing that I would use NVM to first list the available
node versions.

Listing the available node versions:

```bash
nvm ls-remote
```

There I see v18.10.0 is available so I would first try to install that version of node.

Installing a specific version of node:

```bash
nvm install v18.10.0
```

Telling NVM which version of node to use:

```bash
nvm use v18.10.0
```

Setting the default version of node:

```bash
nvm alias default v18.10.0
```

Now that a compatible version of node is installed we can install the Angular CLI globally. Not that I gathered the latest
available version of Angular was 16.2.x so I will install that version of the Angular CLI.  The version numbers may have
changed whenever you read this so please check the [version compatibility](https://angular.io/guide/versions) page to see.

Since your current version of node is v18.10.0 the Angular CLI will be installed globally only for that version of node.
If you change your default version of node to something else you will need to install the Angular CLI globally again.

Installing a version 16.2.x of the Angular CLI:

```bash
npm install -g @angular/cli@16.2.x
```

## References

- [Angular CLI](https://cli.angular.io/)
- [Angular Docs](https://angular.io/docs)
- [Angular Material](https://material.angular.io/)
- [Dockerizing an Angular App](https://mherman.org/blog/dockerizing-an-angular-app/)
- [Angular Version Compatibility](https://angular.io/guide/versions)
- [Node Version Manager](node.md)
