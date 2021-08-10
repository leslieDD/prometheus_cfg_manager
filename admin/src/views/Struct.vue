<template>
  <div>
    <ViewCode :value="structData"></ViewCode>
  </div>
</template>

<script>

import ViewCode from '@/components/ViewCodeMirror.vue'
import { getGoStruct } from '@/api/tmpl.js'

export default {
  name: 'Struct',
  components: {
    ViewCode
  },
  data () {
    return {
      structData: ""
    }
  },
  mounted () {
    this.viewGoStruct()
  },
  methods: {
    viewGoStruct () {
      getGoStruct().then(r => {
        console.log(r.data)
        this.structData = r.data
      }).catch(e => {
        console.log(e)
      })
    }
  }
}
</script>