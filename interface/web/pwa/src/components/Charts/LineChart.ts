import { ChartData, ChartOptions } from "chart.js";
import { Bar, Line } from "vue-chartjs";
import { Component, Mixins, Prop, Vue } from "vue-property-decorator";
import { ChartProperties, DefaultChartOption } from "./ChartInterfaces";

@Component
export default class LineChart extends Mixins(Line) {

  @Prop()
  ChartData:ChartData;
  @Prop()
  ChartOption:ChartOptions;

  public mounted() {
    const opts:ChartOptions = Object.assign(DefaultChartOption ,this.ChartOption);
    console.log(opts);
    this.renderChart(this.ChartData, opts);
  }
}
