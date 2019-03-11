$(function () {


    $('#preTaxIncome').bind('input propertychange', function () {


        initSelect($(this).val());
    });

    $('#select2_group').change(function () {
        var val = $(this).val();
        console.log(val);

    });


    let click = $("#button").click(function (e) {
        const preTaxIncome = $('#preTaxIncome').val();
        const cityCode = $('#select2_group').val();

        const data = JSON.stringify({
            CityCode: cityCode,
            PreTaxIncome: preTaxIncome
        });
        $.ajax({
            type: "POST",
            url: "calc",
            data: data,
            dataType: "json",
            success: function (data) {
                $('#PersonalIncomeTax').html("￥" + data.PersonalIncomeTax);
                $('#PersonalIncomeTaxRate').html(data.PersonalIncomeTaxRate + "%");

                $('#AfterAmount').html("￥" + data.AfterAmount);
                $('#AfterAmountRate').html(data.AfterAmountRate + "%");

                $('#spanAfterAmount').html("￥" + data.AfterAmount);
                $('#styleAfterAmount').attr('style', "width: " + data.AfterAmountRate + "%;");


                $('#Pension').html("￥" + data.Pension);
                $('#PensionRate').html(data.PensionRate + "%");


                $('#Medical').html("￥" + data.Medical);
                $('#MedicalRate').html(data.MedicalRate + "%");


                $('#Provident').html("￥" + data.Provident);
                $('#ProvidentRate').html(data.ProvidentRate + "%");

                $('#Unemployment').html("￥" + data.Unemployment);
                $('#UnemploymentRate').html(data.UnemploymentRate + "%");


                $('#tdQuickDeduction').html("￥" + data.QuickDeduction);
                $('#tdPersonalIncomeTax').html("￥" + data.PersonalIncomeTax);
                $('#tdRate').html(data.Rate + "%");
                $('#rowDetail').html("￥" + data.Amount + "-" + "￥" + data.SocialAmount + "-" + "￥" + data.Exemption);


                $('#spanSocialAmount').html("￥" + data.SocialAmount);


                $('#styleSocialAmount').attr('style', "width: " + data.StyleSocialAmount + "%;");


                $('#spanAmount').html("￥" + data.Amount);

                $('#styleAmount').attr('style', "width: 100%;");


                init_chart_doughnut(data.ProvidentRate, data.UnemploymentRate, data.MedicalRate, data.PensionRate, data.AfterAmountRate, data.PersonalIncomeTaxRate);


            }
        });
    });

    function initSelect(a) {
        const cityCode = $('#select2_group').val();

        $.ajax({
            type: "get",
            url: "getInsuranceByCode/" + cityCode,
            dataType: "json",
            success: function (data) {
                let datum = data[0];
                console.log(datum.PensionUpper);
                //如果输入的值大于等于 社保上限
                if (a >= datum.PensionUpper) {
                    $('#pensionBase').val(datum.PensionUpper);
                } else {
                    if (a > datum.PensionLower) {
                        $('#pensionBase').val(a);
                    } else {
                        $('#pensionBase').val(datum.PensionLower);
                    }
                }

                if (a >= datum.ProvidentUpper) {
                    $('#providentBase').val(datum.ProvidentUpper);
                } else {
                    if (a > datum.ProvidentLower) {
                        $('#providentBase').val(a);
                    } else {
                        $('#providentBase').val(datum.ProvidentLower);
                    }
                }
            }
        });
    }


    function init_chart_doughnut(
        a,
        b,
        c,
        d,
        e,
        f
    ) {
        if (typeof (Chart) === 'undefined') {
            return;
        }
        if (jQuery('.canvasDoughnut').length) {

            const chart_doughnut_settings = {
                type: 'doughnut',
                tooltipFillColor: "rgba(51, 51, 51, 0.55)",
                data: {
                    labels: [
                        "住房公积金",
                        "失业保险金",
                        "医疗保险金",
                        "养老保险金",
                        "税后工资剩余",
                        "个人所得税扣除"
                    ],
                    datasets: [{
                        data: [a, b, c, d, e, f],
                        backgroundColor: [
                            "#BDC3C7", //住房公积金
                            "#9B59B6", //失业保险金
                            "#E74C3C", //医疗保险金
                            "#26B99A", //养老保险金
                            "#3498DB",//税后工资
                            "#F7F709", //个人所得税

                        ],
                        hoverBackgroundColor: [
                            "#CFD4D8",
                            "#B370CF",
                            "#E95E4F",
                            "#36CAAB",
                            "#49A9EA",
                            "#F7F709",


                        ]
                    }]
                },
                options: {
                    legend: false,
                    responsive: false
                }
            };

            $('.canvasDoughnut').each(function () {

                var chart_element = $(this);
                var chart_doughnut = new Chart(chart_element, chart_doughnut_settings);

            });

        }

    }

})
;