# fedramp-ksi-report

A Go CLI that takes YAML input and produces a Markdown document representing the [Continuous Monitoring Monthly key security indicator report](https://github.com/FedRAMP/rev5-continuous-monitoring-cwg/blob/main/conmon_ksi_report.md) proposed by the [FedRAMP 20x Continuous Monitoring Working Group](https://github.com/FedRAMP/rev5-continuous-monitoring-cwg/tree/main)

## Usage

The CLI produced by this project is named `fkr` and it produces the sample report in `sample_report.md` when running

```shell
./fkr > sample_report.md
```

which is equivalent to running

```shell
./fkr --input ./ksi.yaml --template ./ksi_report.md.tmpl > sample_report.md
```

## Output

See [sample_report.md](sample_report.md) for an example of the rendered report with sample data.