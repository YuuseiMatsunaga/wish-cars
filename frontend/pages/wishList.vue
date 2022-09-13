<template>
    <div >
        <v-row v-for="row of displayRowCount" :key="row">
            <v-row justify="space-between" class="mt-5">
                <div v-for="n of 2" :key="n">
                    <v-row align="center">
                        <img :src="cars[row*2 + n - 3].maker.logo_path" alt="" width="50" height="50"/>
                        <div style="font-size: 40px">{{cars[row*2 + n - 3].name}}</div>
                    </v-row>
                    <img class="mt-3" :src="cars[row*2 + n - 3].img_path" width="580" height="280"/>
                </div>
            </v-row>
        </v-row>
    </div>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data () {
    return {
        cars: [
            {
                id: 1,
                maker: { logo_path: '/public/maker-logos/alfaromeo.png' },
                name: '4c',
                img_path: '/public/car-images/alfa4c.png'
            },
            {
                id: 2,
                maker: { logo_path: '/public/maker-logos/dadge.png' },
                name: 'Challenger',
                img_path: '/public/car-images/dadgechallenger.png'
            },
            {
                id: 3,
                maker: { logo_path: '/public/maker-logos/ferrari.png' },
                name: 'Lusso',
                img_path: '/public/car-images/gt4lusso.png'
            },
            {
                id: 4,
                maker: { logo_path: '/public/maker-logos/lotus.png' },
                name: 'Emira',
                img_path: '/public/car-images/lotusemira.png'
            }
        ]
    }
  },
  computed:{
    displayRowCount() {
        return Math.ceil(this.cars.length/2)
    }
  }
})
</script>