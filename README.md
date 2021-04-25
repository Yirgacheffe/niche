Learning & demo purpose repository - Go based services
------------------------------

Successfully tested on Docker-Desktop Kubernetes and Istio 1.6.7.

## Prerequisites

You're going to need:
 - **Kubernetes**
 - **Istio**
 - **Go**
 - **Docker**
 - **Amqp** (CloudAmqp recommanded)
 - **MongoDB**

## How to run it locally

- Make sure you have Kubernetes, Docker and Istio installed and running on your machine.
- Run 'build_images.sh' in 'configs' and 'web' directory
- Using manifest files in 'deployment' directory to setup
- Using menifect files in 'deployment/configs' and 'deployment/web' to deploy the service 'configs' and 'web'

- Module 'greetings' is to demostrate 'Service Tracing with Jeager' and 'Kiali' features.
- Module 'greetings' need MongoDB and CloudAmqp, please make sure you already have it.


Note***: 'todox', 'tus' and 'authx' for practice purpose, will try to complete it. Have Fun!!!