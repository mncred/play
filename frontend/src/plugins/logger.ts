import * as backend from '$app/logrus/Logger'

// builds single string for logger
const render = (...data: any[]) => {
    if (data.length === 0) return ''
    const message = typeof data[0] === 'string' ? data[0] : JSON.stringify(data[0])
    const args = data.slice(1).length > 0 ? JSON.stringify(data.slice(1)) : ''
    return `${message} ${args}`
}

/**
 * Plugin which binds default browser logging mechanism with backend logging.
 * It allows to handle all frontend logs from backend.
 * For frontend developers it looks like native console logging usage.
 */
const pluginLogger = async () => {
    const logger = {
        install: () => {
            // save refs to default browser log functions
            const frontend = { ...console }
            // replace most common default browser log functions
            window.console = {
                ...window.console,
                log: (...data: any[]) => {
                    frontend.log(...data)
                    backend.Info(render(...data))
                },
                debug: (...data: any[]) => {
                    frontend.debug(...data)
                    backend.Debug(render(...data))
                },
                info: (...data: any[]) => {
                    frontend.info(...data)
                    backend.Info(render(...data))
                },
                warn: (...data: any[]) => {
                    frontend.warn(...data)
                    backend.Warning(render(...data))
                },
                error: (...data: any[]) => {
                    frontend.error(...data)
                    backend.Error(render(...data))
                },
            }
            console.info('pluginLogger installed, all frontend logs will be mirrored to backend')
        }
    }
    return logger
}

export default pluginLogger