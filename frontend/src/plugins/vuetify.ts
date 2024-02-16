import 'vuetify/styles'
import '@fontsource/roboto/latin.css'
import '@fontsource/roboto/latin-italic.css'
import '@fontsource/roboto/cyrillic.css'
import '@fontsource/roboto/cyrillic-italic.css'

import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import { aliases, mdi } from 'vuetify/iconsets/mdi-svg'

import { createVuetify } from 'vuetify'

const pluginVuetify = async () => {
    const vuetify = createVuetify({
        components,
        directives,
        icons: {
            defaultSet: 'mdi',
            aliases,
            sets: {
                mdi
            }
        }
    })

    return vuetify
}

export default pluginVuetify