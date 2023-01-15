import { sleep } from 'k6';
import { generateModel, generatePipelineName } from '../components/model.js';
import { connectScheduler, disconnectScheduler, loadModel, unloadModel, loadPipeline, unloadPipeline, awaitPipelineStatus } from '../components/scheduler.js';
import {
    connectScheduler as connectSchedulerProxy,
    disconnectScheduler as disconnectSchedulerProxy,
    loadModel as loadModelProxy,
    unloadModel as unloadModelProxy
} from '../components/scheduler_proxy.js';
import { inferGrpcLoop, inferHttpLoop, modelStatusHttp } from '../components/v2.js';

export function setupBase(config ) {
    if (config.loadModel) {
        
        var connectSchedulerFn = connectScheduler
        if (config.isSchedulerProxy) {
            connectSchedulerFn = connectSchedulerProxy
        }
        
        connectSchedulerFn(config.schedulerEndpoint)

        for (let i = 0; i < config.maxNumModels; i++) {
            const modelName = config.modelNamePrefix + i.toString()
            const model = generateModel(config.modelType, modelName, 1, 1, config.isSchedulerProxy, config.modelMemoryBytes, config.inferBatchSize)
            const modelDefn = model.modelDefn
            const pipelineDefn = model.pipelineDefn

            var loadModelFn = loadModel
            if (config.isSchedulerProxy) {
                loadModelFn = loadModelProxy
            }

            loadModelFn(modelName, modelDefn, false)

            if (config.isLoadPipeline) {
                loadPipeline(generatePipelineName(modelName), pipelineDefn, false)  // we use pipeline name as model name
            }
        }

        // note: this doesnt work in case of kafka
        // and in the pipeline gateway doesnt support status endpoint
        // wait for all models to be ready
        if (!config.isLoadPipeline) {
            const n = config.maxNumModels - 1
            const modelName = config.modelNamePrefix + n.toString()
            const modelNameWithVersion = modelName + getVersionSuffix(config)  // first version
            while (modelStatusHttp(config.inferHttpEndpoint, config.isEnvoy?modelName:modelNameWithVersion, config.isEnvoy) !== 200) {
                sleep(1)
            }
        } else {
            const n = config.maxNumModels - 1
            const modelName = config.modelNamePrefix + n.toString()
            awaitPipelineStatus(generatePipelineName(modelName), "PipelineReady")
        }

        if (config.doWarmup) {
            // warm up
            for (let i = 0; i < config.maxNumModels; i++) {
                const modelName = config.modelNamePrefix + i.toString()

                const modelNameWithVersion = modelName + getVersionSuffix(config)  // first version
                
                const model = generateModel(config.modelType, modelNameWithVersion, 1, 1, config.isSchedulerProxy, config.modelMemoryBytes)

                inferHttpLoop(
                    config.inferHttpEndpoint, config.isEnvoy?modelName:modelNameWithVersion, model.inference.http, 1, config.isEnvoy, config.dataflowTag)
            }
        }

        var disconnectSchedulerFn = disconnectScheduler
        if (config.isSchedulerProxy) {
            disconnectSchedulerFn = disconnectSchedulerProxy
        }
        disconnectSchedulerFn()
    }
}

export function teardownBase(config ) {
    if (config.unloadModel) {
        var connectSchedulerFn = connectScheduler
        if (config.isSchedulerProxy) {
            connectSchedulerFn = connectSchedulerProxy
        }
        connectSchedulerFn(config.schedulerEndpoint)

        var unloadModelFn = unloadModel
        if (config.isSchedulerProxy) {
            unloadModelFn = unloadModelProxy
        }

        for (let i = 0; i < config.maxNumModels; i++) {
            const modelName = config.modelNamePrefix + i.toString()
            // if we have added a pipeline, unloaded it
            if (config.isLoadPipeline) {
                unloadPipeline(generatePipelineName(modelName)) 
            }

            unloadModelFn(modelName, false)
        }

        var disconnectSchedulerFn = disconnectScheduler
        if (config.isSchedulerProxy) {
            disconnectSchedulerFn = disconnectSchedulerProxy
        }
        disconnectSchedulerFn()
    }
}

export function doInfer(modelName, modelNameWithVersion, config, isHttp) {
    const model = generateModel(config.modelType, modelName, 1, 1, config.isSchedulerProxy, config.modelMemoryBytes, config.inferBatchSize)
    const httpEndpoint = config.inferHttpEndpoint
    const grpcEndpoint = config.inferGrpcEndpoint

    if (config.infer) {
        if (isHttp) {
            inferHttpLoop(
                httpEndpoint, config.isEnvoy?modelName:modelNameWithVersion, model.inference.http, config.inferHttpIterations, config.isEnvoy, config.dataflowTag)
        } else {
            inferGrpcLoop(
                grpcEndpoint, config.isEnvoy?modelName:modelNameWithVersion, model.inference.grpc, config.inferGrpcIterations, config.isEnvoy, config.dataflowTag)
        }
    }
}

export function getVersionSuffix(config) {
    var versionSuffix = "_1"
    if (config.isSchedulerProxy) {
        versionSuffix = "_0"
    }
    return versionSuffix
}