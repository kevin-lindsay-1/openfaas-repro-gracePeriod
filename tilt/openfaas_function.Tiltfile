load(
    './helmfile.Tiltfile',
    helmfile='helmfile'
)

# Function
k8s_kind(
    'Function',
    api_version='openfaas.com/.*',
    image_json_path="{ .spec.image }"
)


def openfaas_function(
    function_name,
    template_name,
    environment='local',
    function_directory='.',
    k8s_directory='k8s',
    k8s_out_directory='k8s.out',
    namespace='openfaas-fn',
    prefix=None,
    shrinkwrap_name='shrinkwrapped',
    build_args={},
):
    watch_file(k8s_directory)

    helmfile_path = '{0}/helmfile.yaml'.format(k8s_directory)
    template_directory = 'template/{0}'.format(
        template_name,
    )
    shrinkwrap_directory = 'build/{0}'.format(
        shrinkwrap_name,
    )

    shrinkwrap_cmd = ''' \
    faas build --shrinkwrap --image={1} --name={1} --lang={2} --handler={0} \
    && \
    mv build/{1}/function/.dockerignore build/{1}/.dockerignore \
    '''.format(
        function_directory,
        shrinkwrap_name,
        template_name,
    )

    k8s_resource_deps = []

    if environment == 'local':
        if not os.path.exists(shrinkwrap_directory):
            local(shrinkwrap_cmd)

        shrinkwrap_resource_name = '{0}-shrinkwrap'.format(
            function_name,
        )
        if prefix:
            shrinkwrap_resource_name = '{0}-{1}-shrinkwrap'.format(
                prefix,
                function_name,
            )
        local_resource(
            shrinkwrap_resource_name,
            cmd=shrinkwrap_cmd,
            deps=[
                function_directory,
            ],
            ignore=[
                k8s_directory,
                k8s_out_directory,
                shrinkwrap_directory,
            ],
        )
        k8s_resource_deps += [shrinkwrap_resource_name]

        if os.path.exists(shrinkwrap_directory):
            docker_build(
                function_name,
                shrinkwrap_directory,
                build_args,
            )

    k8s_yaml(
        helmfile(
            function_name,
            helmfile_path=helmfile_path,
            namespace=namespace,
            environment=environment,
        )
    )

    function_resource_name = function_name
    if prefix:
        function_resource_name = '{0}-{1}'.format(prefix, function_name)

    k8s_resource(
        function_name,
        new_name=function_resource_name,
        resource_deps=k8s_resource_deps,
        extra_pod_selectors=[
            {
                'faas_function': function_name,
            },
        ],
    )
