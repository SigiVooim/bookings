function Prompt() {
    let toast = function (c) {
        const {
            msg = '',
            icon = 'success',
            position = 'top-end',

        } = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }

    let success = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c

        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer,
        })

    }

    let error = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c

        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
        })

    }

    async function custom(c) {
        const {
            msg = "",
            title = "",
        } = c;

        const {value: formValues} = await Swal.fire({
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () => {
              const elem = document.getElementById("reservation-dates");
              const rp = new DateRangePicker(elem, {
                  format: 'yyyy-mm-dd',
                  showOnFocus: true,
              })
            },
            preConfirm: () => {
                return [
                    document.getElementById("start").value,
                    document.getElementById("end").value
                ];
            }
        });
        if (formValues) {
            Swal.fire(JSON.stringify(formValues));
        }

    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}