import { ChartOptions } from "chart.js"

export interface ChartProperties {
    Data: any[],
    XAxisLabels: string[],
    ChartOptions: any
}

export const DefaultChartOption: ChartOptions = {
    legend: {
        position: 'bottom',
    },
    title: {
        display: true,
        text: 'Default Chart Title'
    },
    maintainAspectRatio: false
}