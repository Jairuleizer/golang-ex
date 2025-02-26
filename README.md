Golang Sample App on OpenShift
==============================

This is a sample Golang web application for OpenShift v3 that use the [beego](http://beego.me/)
framework. This sample uses official Docker [golang](https://hub.docker.com/_/golang/) image for the build
without any modification to the image or sample itself.

This example was copied from the official [beego samples](https://github.com/beego/samples) repository.

If you'd like to install it, follow [these instructions](https://github.com/Jairuleizer/golang-ex/blob/master/README.md#installation).

The steps in this document assume that you have access to an OpenShift deployment that you can deploy applications on, including the default set of ImageStreams defined. Instructions for installing the default ImageStreams are available [here](https://docs.openshift.org/latest/install_config/imagestreams_templates.html). If you are defining the set of ImageStreams now, remember to pass in the proper cluster-admin credentials and to create the ImageStreams in the 'openshift' namespace.

### Installation

1. Fork a copy of [golang-ex](https://github.com/Jairuleizer/golang-ex)
2. Clone your repository to your development machine and cd to the repository directory
3. Add a Golang application from the beego template. **Note:** You will need to replace part of the `SOURCE_REPOSITORY_URL` where it says `yourusername` with your GitHub username before running:

        $ oc new-app openshift/templates/beego.yml -p SOURCE_REPOSITORY_URL=https://github.com/yourusername/golang-ex

4. A build will be started automatically. It might take some time.  
You can run the command below to watch for builds:

        $ oc get builds -w

5. Once the build is running, follow the build logs:

        $ oc logs -f bc/beego-fatesg

6. Wait for the beego-fatesg pod to start up (may take some time):

        $ oc get pods -w

    Sample output:

        NAME                     READY     REASON       RESTARTS   AGE
        beego-fatesg-1-6c23l     1/1       Running        0          2m
        beego-fatesg-1-build     0/1       ExitCode:0     0          4m

7. Check the IP and port the beego-fatesg service is running on:

        $ oc get svc

    Sample output:

        NAME             LABELS                              SELECTOR              IP(S)           PORT(S)
        beego-fatesg     template=beego-fatesg     name=beego-fatesg           172.30.210.29    8080/TCP

In this case, the IP for beego-fatesg is 172.30.210.29 and it is on port 8080.  
**Note**: you can also get this information from the web console.

### Building

After you create the application from a template, you can follow the build progress
with the `oc logs` command:

```console
$ oc logs -f bc/beego-fatesg
```

### Accessing the application

If you have the OpenShift router running, you should be able to access the
application just by typing the beego-fatesg route DNS
(beego-fatesg.openshiftapps.com) into your browser. However, you will have to
run you own DNS server first, but we can cheat it by modifying the `/etc/hosts`
file on your host machine. Just append this line at the end of that file:

```
192.168.124.206 beego-fatesg.openshiftapps.com
```

The `192.168.124.206` represents the IP address of the machine you are running
OpenShift on. If your are running OpenShift on a different IP address, simply replace
`192.168.124.206` with your OpenShift IP address when adding the line above
to your `/etc/hosts` file. Once you have this change, you can just type the
application DNS into browser and you should see the chat application login screen.

### License

This code is dedicated to the public domain to the maximum extent permitted by applicable law, pursuant to [CC0](http://creativecommons.org/publicdomain/zero/1.0/).

The sample beego application code follows the original [Apache License](https://github.com/beego/samples/blob/master/LICENSE)
