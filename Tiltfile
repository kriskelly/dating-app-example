load('ext://restart_process', 'docker_build_with_restart')

k8s_yaml('deployments/api.yaml')

docker_build_with_restart('dating-app/api', '.',
    entrypoint='/app/build/api',
    ignore=['./Dockerfile', '.git'],
    live_update=[
        sync('.', '/app'),
        run('go build -o /app/build/api /app/cmd'),
    ]
)

k8s_resource('api', port_forwards=[3000])