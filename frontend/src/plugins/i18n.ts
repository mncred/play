import { createI18n } from 'vue-i18n'

const pluginI18n = async () => {
    // object with locales map for i18n
    const locales = {} as Record<string, any>

    // iterate over all available locales inside locales dir
    const files = await import.meta.glob('@/locales/*.yaml')
    for (const filename in files) {
        const code = filename.match(/\/([^\/]+)\.ya?ml$/)?.[1]
        if (code) {
            const translations = (await files[filename]() as any).default
            locales[code] = translations
        }
    }

    // create Vue plugin
    const i18n = createI18n({
        globalInjection: true,
        locale: 'ru-RU',
        fallbackLocale: 'ru-RU',
        messages: locales,
    })

    return i18n
}

export default pluginI18n