module.exports = {
    branches: ['main'],
    plugins: [
        '@semantic-release/commit-analyzer',
        '@semantic-release/release-notes-generator',
        '@semantic-release/changelog',
        [
            '@semantic-release/exec',
            {
                prepareCmd: 'echo ${nextRelease.version} > VERSION',
            },
        ],
        '@semantic-release/npm',
        '@semantic-release/github',
    ],
};
