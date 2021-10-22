###
# CONSTANTS
###
PROJECT_NAME = 'sleep'
FUNCTION_TEMPLATE = 'golang-http'

###
# LOAD HELPERS
###
load(
    './tilt/config.Tiltfile',
    get_config='get_config'
)
load(
    './tilt/helmfile.Tiltfile',
    helmfile='helmfile'
)
load(
    './tilt/openfaas_function.Tiltfile',
    openfaas_function='openfaas_function'
)

###
# INITIALIZATION
###
# Get project config
cfg = get_config()
project_config = cfg.get(PROJECT_NAME, {})

# Determine whether or not to use the root directory for dependencies
# (useful if you have two different branches for a service)
use_main_dir_for_dependencies = project_config.get(
    'use_main_dir_for_dependencies',
    True
)

# Specify a k8s namespace in tilt_config based off different contexts
# (certain apps may be in more than one namespace)
namespace = project_config.get('namespace', 'default')

# Dynamically include Tiltfiles
dependencies = project_config.get('dependencies', [])
dependency_names = [dependency.get('name') for dependency in dependencies]
dependency_tiltfile_paths = [
    dependency.get('tiltfile_path') for dependency in dependencies
]

for tiltfile_path in dependency_tiltfile_paths:
    if tiltfile_path != None:
        if use_main_dir_for_dependencies:
            include('{0}/{1}'.format(config.main_dir, tiltfile_path))
        else:
            include(tiltfile_path)

# Get list of resources to skip
skip_resources = project_config.get('skip_resources', [])

# Get optional environment variant
environment = project_config.get('environment', 'local')

openfaas_function(
    PROJECT_NAME,
    FUNCTION_TEMPLATE,
    environment=environment,
)
