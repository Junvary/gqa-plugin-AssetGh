<template>
    <q-page padding>
        <div class="row q-gutter-md items-center justify-between" style="margin-bottom: 10px">
            <q-btn color="primary" @click="resetSettlement" :label="'重新进行' + nowYearMonth + '月结'" />
            <div class="row items-center q-gutter-sm">
                <q-input v-model="queryParams.setYearMonth" label="选择年月" clearable placeholder="202201" />
                <q-btn color="primary" @click="getTableDataTitle" label="查询" />
                <q-btn color="primary" @click="resetYearMonth" label="重置" />
                <q-btn color="primary" icon="print" v-print="'#printContent'">打印</q-btn>
                <q-btn color="primary" icon-right="archive" label="导出" no-caps @click="exportTable" />
            </div>

        </div>
        <q-table id="printContent" row-key="assetCatalog1" separator="cell" :rows="tableData" :columns="columns"
            hide-bottom v-model:pagination="pagination" :rows-per-page-options="pageOptions" :loading="loading"
            @request="onRequest" :title="'固定资产计提折旧（' + titleYearMonth + '）'">

            <template v-slot:body-cell-monthDepreciation="props">
                <q-td :props="props">
                    {{ props.row.monthDepreciation.toFixed(2) }}
                </q-td>
            </template>

            <template v-slot:body-cell-totalDepreciation="props">
                <q-td :props="props">
                    {{ props.row.totalDepreciation.toFixed(2) }}
                </q-td>
            </template>

            <template v-slot:body-cell-netWorth="props">
                <q-td :props="props">
                    {{ props.row.netWorth.toFixed(2) }}
                </q-td>
            </template>

            <!-- <template v-slot:top="props">
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template> -->

        </q-table>
    </q-page>
</template>


<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { date, exportFile } from 'quasar'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-AssetGh/get-settlement-list',
    set: 'plugin-AssetGh/set-settlement',
}
const columns = computed(() => {
    return [
        { name: 'setYearMonth', align: 'center', label: '月结', field: 'setYearMonth' },
        { name: 'assetCatalog1', align: 'center', label: '固定资产类别', field: 'assetCatalog1' },
        { name: 'originalValue', align: 'center', label: '固定资产原值', field: 'originalValue' },
        { name: 'monthDepreciation', align: 'center', label: '月累计折旧', field: 'monthDepreciation' },
        { name: 'totalDepreciation', align: 'center', label: '累计折旧额', field: 'totalDepreciation' },
        { name: 'netWorth', align: 'center', label: '固定资产净值', field: 'netWorth' },
    ]
})
const {
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
    const timeStamp = Date.now()
    queryParams.value.setYearMonth = date.formatDate(timeStamp, 'YYYYMM')
    nowYearMonth.value = date.formatDate(timeStamp, 'YYYYMM')
    titleYearMonth.value = nowYearMonth.value
    getTableData()
})
const nowYearMonth = ref('')
const titleYearMonth = ref('')
const resetSettlement = () => {
    postAction(url.set).then((res) => {
        if (res.code === 1) {
            $q.notify({
                type: 'positive',
                message: res.message,
            })
            getTableData()
        }
    })
}
const resetYearMonth = () => {
    const timeStamp = Date.now()
    queryParams.value.setYearMonth = date.formatDate(timeStamp, 'YYYYMM')
    titleYearMonth.value = nowYearMonth.value
    getTableData()
}
const getTableDataTitle = () => {
    titleYearMonth.value = queryParams.value.setYearMonth
    getTableData()
}
const exportTable = () => {
    // naive encoding to csv format
    const content = [columns.value.map((col) => wrapCsvValue(col.label))].concat(tableData.value.map((row) => columns.value.map((col) => wrapCsvValue(typeof col.field === 'function' ? col.field(row) : row[col.field === void 0 ? col.name : col.field], col.format)).join(','))).join('\r\n')

    const status = exportFile('固定资产计提折旧(' + titleYearMonth.value + ').csv', "\ufeff" + content, 'text/csv')

    if (status === true) {
        $q.notify({
            type: 'positive',
            message: '开始文件导出！',
        })
    } else {
        $q.notify({
            type: 'negative',
            message: '文件导出失败！',
        })
    }

}
const wrapCsvValue = (val, formatFn) => {
    let formatted = formatFn !== void 0 ? formatFn(val) : val

    formatted = formatted === void 0 || formatted === null ? '' : String(formatted)

    formatted = formatted.split('"').join('""')
    /**
     * Excel accepts \n and \r in strings, but some other CSV parsers do not
     * Uncomment the next two lines to escape new lines
     */
    // .split('\n').join('\\n')
    // .split('\r').join('\\r')

    return `"${formatted}"`
}
</script>
