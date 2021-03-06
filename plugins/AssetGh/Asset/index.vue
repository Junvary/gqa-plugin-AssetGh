<template>
    <q-page padding>

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 15%" v-model="queryParams.assetCode" label="资产编号" />
            <q-input style="width: 15%" v-model="queryParams.assetName" label="资产名称（模糊）" />
            <q-select style="width: 15%" v-model="queryParams.assetCatalog1" :options="dictOptions.AssetGhCatalog"
                clearable emit-value map-options label="资产大类" />
            <q-select style="width: 15%" v-model="queryParams.assetCatalog2" :options="catalog2" clearable emit-value
                map-options label="资产细类" />
            <q-select style="width: 15%" v-model="queryParams.useStatus" :options="dictOptions.AssetGhUseStatus"
                clearable emit-value map-options label="使用状态" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增工会固定资产" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-alreadyUse="props">
                <q-td :props="props">
                    {{ diffDateFromEntryDate(props.row) }}
                </q-td>
            </template>

            <template v-slot:body-cell-depreciationMonth="props">
                <q-td :props="props">
                    {{ props.row.depreciationMonth.toFixed(2) }}
                </q-td>
            </template>

            <template v-slot:body-cell-alreadyDepreciation="props">
                <q-td :props="props">
                    {{ getAlreadyDepreciation(props.row) }}
                </q-td>
            </template>

            <template v-slot:body-cell-useStatus="props">
                <q-td :props="props">
                    <GqaDictShow dictName="AssetGhUseStatus" :dictCode="props.row.useStatus" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>


<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail'
import { date } from 'quasar'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-AssetGh/get-asset-list',
    delete: 'plugin-AssetGh/delete-asset-by-id',
}
const columns = computed(() => {
    return [
        { name: 'assetCode', align: 'center', label: '资产编号', field: 'assetCode' },
        { name: 'assetName', align: 'center', label: '资产名称', field: 'assetName' },
        { name: 'entryDate', align: 'center', label: '入账日期', field: 'entryDate' },
        { name: 'usefulLife', align: 'center', label: '预计使用年限', field: 'usefulLife' },
        { name: 'alreadyUse', align: 'center', label: '已使用(月)', field: 'alreadyUse' },
        { name: 'number', align: 'center', label: '数量', field: 'number' },
        { name: 'originalValue', align: 'center', label: '资产原值(元)', field: 'originalValue' },
        { name: 'depreciationMonth', align: 'center', label: '计提折旧(元月)', field: 'depreciationMonth' },
        { name: 'alreadyDepreciation', align: 'center', label: '累计折旧(元)', field: 'alreadyDepreciation' },
        { name: 'useStatus', align: 'center', label: '使用状态', field: 'useStatus' },
        { name: 'actions', align: 'center', label: '操作', field: 'actions' },
    ]
})
const {
    dictOptions,
    showDateTime,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    pagination.value.sortBy = 'entry_date'
    pagination.value.descending = true
    getTableData()
})
const catalog2 = computed(() => {
    if (queryParams.value.assetCatalog1) {
        return dictOptions.value.AssetGhCatalog.find((item) => item.value === queryParams.value.assetCatalog1).children
    }
})
const diffDateFromEntryDate = computed(() => {
    return (row) => {
        if (row) {
            let now = new Date()
            let entryDate = new Date(row.entryDate)
            let diffMonth = date.getDateDiff(now, entryDate, 'month')

            let scrapDate = new Date(row.scrapDate)
            const diff = date.getDateDiff(now, scrapDate)
            if (diff > 0) {
                return '-'
            } else {
                return diffMonth
            }
        }
    }
})
const getAlreadyDepreciation = computed(() => {
    return (row) => {
        if (row.entryDate) {
            let now = new Date()
            let scrapDate = new Date(row.scrapDate)
            const diff = date.getDateDiff(now, scrapDate)
            const nowDate = date.formatDate(now, 'YYYY-MM')
            if (row.scrapDate.slice(0, 7) === nowDate.slice(0, 7)) {
                return row.originalValue * row.number
            } else if (diff > 0) {
                return row.originalValue * row.number
            } else {
                let entryDate = new Date(row.entryDate)
                let diffMonth = date.getDateDiff(now, entryDate, 'month')
                console.log(diffMonth)
                const depreciationMonth = row.depreciationMonth
                const alreadyDepreciation = depreciationMonth * diffMonth
                return alreadyDepreciation
            }
        }
    }
})
</script>
