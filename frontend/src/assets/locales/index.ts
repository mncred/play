import { createI18n } from 'vue-i18n'
import locale_en_US from './en-US.json'
import locale_ru_RU from './en-US.json'

const i18n = createI18n({
    locale: 'ru-RU',
    fallbackLocale: 'en-US',
    messages: {
        'en-US': locale_en_US,
        'ru-RU': locale_ru_RU,
    }
})

export default i18n