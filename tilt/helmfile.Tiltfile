"""Helmfile-related helpers"""


def helmfile(
    release_name,
    helmfile_path='./k8s/helmfile.yaml',
    release_version='1.0.0',
    # relative to helmfile
    release_chart='./charts/app',
    namespace='default',
    environment='local'
):
    """Helper function to read K8s config YAML from helmfile."""

    watch_file(helmfile_path)

    # `local`'s `cwd` is relative to tiltfile
    local('rm -rf ./k8s.out')

    # `helmfile`'s `cwd` is relative to the `helmfile.yaml`
    local(
        ''.join(
            (
                'RELEASE_NAME={0} ',
                'RELEASE_VERSION={1} ',
                'RELEASE_CHART={2} ',
                'helmfile',
                ' --debug',
                ' --file={3}',
                ' --namespace={4}',
                ' --environment={5}',
                ' template',
                ' --output-dir=../k8s.out'
            )
        )
        .format(
            release_name,
            release_version,
            release_chart,
            helmfile_path,
            namespace,
            environment
        )
    )
    return local(
        ''.join(
            (
                'RELEASE_NAME={0} ',
                'RELEASE_VERSION={1} ',
                'RELEASE_CHART={2} ',
                'helmfile',
                ' --file={3}',
                ' --namespace={4}',
                ' --environment={5}',
                ' template',
            )).format(
            release_name,
            release_version,
            release_chart,
            helmfile_path,
            namespace,
            environment,
        )
    )
