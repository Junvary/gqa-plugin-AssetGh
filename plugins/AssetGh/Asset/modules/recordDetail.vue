<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }}工会固定资产:
                    {{ recordDetail.value.title }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.assetCode" label="资产编号"
                            :rules="[val => val && val.length > 0 || '必须输入资产编号']" />
                        <q-input class="col" v-model="recordDetail.value.assetName" label="资产名称"
                            :rules="[val => val && val.length > 0 || '必须输入资产名称']" />
                    </div>
                    <div class="row">
                        <q-select class="col" v-model="recordDetail.value.assetCatalog1"
                            @update:model-value="changeCatalog1" :options="dictOptions.AssetGhCatalog" clearable
                            emit-value map-options :rules="[val => val && val.length > 0 || '必须选择资产大类']" label="资产大类" />
                        <q-select class="col" v-model="recordDetail.value.assetCatalog2" :options="catalog2" clearable
                            emit-value map-options :rules="[val => val && val.length > 0 || '必须选择资产细类']" label="资产细类"
                            @update:model-value="changeCatalog2" />
                        <q-input class="col" v-model.number="recordDetail.value.usefulLife" type="number" label="使用年限"
                            :rules="[val => checkUsefulLife(val)]" />
                    </div>
                    <div class="row">
                        <q-input class="col" v-model.number="recordDetail.value.originalValue" label="资产原值(元)"
                            type="number" :rules="[val => val && val > 0 || '必须输入资产原值,资产原值必须大于0']" />
                        <q-input class="col" stack-label v-model="recordDetail.value.entryDate" label="入账日期" type="date"
                            :rules="[val => val && val.length > 0 || '必须输入入账日期']" />
                        <q-input class="col" v-model.number="recordDetail.value.number" label="数量" type="number"
                            :rules="[val => val && val > 0 || '必须输入数量,数量必须大于0']" />
                    </div>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.userDept" label="使用部门"
                            :rules="[val => val && val.length > 0 || '必须输入使用部门']" />
                        <q-input class="col" v-model="recordDetail.value.storageLocation" label="存放地点"
                            :rules="[val => val && val.length > 0 || '必须输入存放地点']" />
                        <q-input class="col" v-model="recordDetail.value.custodian" label="保管人"
                            :rules="[val => val && val.length > 0 || '必须输入保管人']" />
                        <q-field dense label="使用状态" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.useStatus"
                                    :options="dictOptions.AssetGhUseStatus" color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                    </div>
                    <div class="row ">
                        <q-input class="col" v-model="recordDetail.value.memo" type="textarea" label="备注" />
                    </div>

                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="$t('Save') + formTypeName" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('Cancel')" color="negative" v-close-popup />
            </q-card-actions>

            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>


<script setup>
import useRecordDetail from 'src/composables/useRecordDetail'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaShowName from 'src/components/GqaShowName'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { ref, computed, watch } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import { date } from 'quasar'

const $q = useQuasar()
const { t } = useI18n()
const storageStore = useStorageStore()
const gqaBackend = computed(() => storageStore.GetGqaBackend())
const emit = defineEmits(['handleFinish'])
const url = {
    add: 'plugin-AssetGh/add-asset',
    edit: 'plugin-AssetGh/edit-asset',
    queryById: 'plugin-AssetGh/query-asset-by-id',
}
const {
    dictOptions,
    showDateTime,
    formType,
    formTypeName,
    recordDetail,
    recordDetailVisible,
    loading,
    // show,
    recordDetailForm,
    handleAddOrEidt
} = useRecordDetail(url, emit)

const usefulLife1 = ref('')
const usefulLife2 = ref('')

const show = async (row) => {
    loading.value = true
    recordDetail.value = Object.assign({}, row)
    recordDetailVisible.value = true
    if (!recordDetail.value.id) {
        loading.value = false
    } else {
        await handleQueryById(recordDetail.value.id)
    }
    let now = new Date()
    let entryDate = new Date(recordDetail.value.entryDate)
    let diffMonth = date.getDateDiff(now, entryDate, 'month')
    if (diffMonth > recordDetail.value.usefulLife * 12) {
        $q.notify({
            type: 'warning',
            message: '使用年限已超，判定为报废！',
        })
        recordDetail.value.useStatus = 'assetGh_scrap'
    } else {
        recordDetail.value.useStatus = 'assetGh_on_use'
    }
}
const handleQueryById = async (id) => {
    postAction(url.queryById, {
        id,
    }).then(res => {
        if (res.code === 1) {
            recordDetail.value = res.data.records
        }
    }).finally(() => {
        loading.value = false
    })
}

defineExpose({
    show,
    formType
})
const catalog2 = computed(() => {
    if (recordDetail.value.assetCatalog1) {
        return dictOptions.value.AssetGhCatalog.find((item) => item.value === recordDetail.value.assetCatalog1).children
    }
})
const checkUsefulLife = computed(() => {
    return (val) => {
        return (val && val >= usefulLife1.value && (usefulLife2.value ? val <= usefulLife2.value : true)) || '必须输入使用年限' + (usefulLife1.value ? `，大于${usefulLife1.value}` : '') + (usefulLife2.value ? `，小于${usefulLife2.value}` : '')
    }
})
watch(() => recordDetail.value?.usefulLife, (val) => {
    if (val && recordDetail.value.entryDate) {
        let now = new Date()
        let entryDate = new Date(recordDetail.value.entryDate)
        let diffMonth = date.getDateDiff(now, entryDate, 'month')
        if (diffMonth > val * 12) {
            $q.notify({
                type: 'warning',
                message: '使用年限已超，判定为报废！',
            })
            recordDetail.value.useStatus = 'assetGh_scrap'
        } else {
            recordDetail.value.useStatus = 'assetGh_on_use'
        }
    }
})
watch(() => recordDetail.value?.entryDate, (val) => {
    if (val && recordDetail.value.usefulLife) {
        let now = new Date()
        let entryDate = new Date(val)
        let diffMonth = date.getDateDiff(now, entryDate, 'month')
        if (diffMonth > recordDetail.value.usefulLife * 12) {
            $q.notify({
                type: 'warning',
                message: '使用年限已超，判定为报废！',
            })
            recordDetail.value.useStatus = 'assetGh_scrap'
        } else {
            recordDetail.value.useStatus = 'assetGh_on_use'
        }
    }
})

const changeCatalog1 = () => {
    recordDetail.value.assetCatalog2 = ''
    recordDetail.value.usefulLife = null
}
const changeCatalog2 = (value) => {
    const life = catalog2.value.find((item) => item.value === value)
    usefulLife1.value = life.dict_ext_1
    usefulLife2.value = life.dict_ext_2
    recordDetail.value.usefulLife = null
}

</script>
