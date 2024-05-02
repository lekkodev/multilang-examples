examples_config_v1beta1 = proto.package("examples.config.v1beta1")

export(
    Config(
        description = "tunable interface",
        default = examples_config_v1beta1.TunableStruct(
            string_field = "default",
            number_field = 42,
            boolean_field = True,
        ),
        overrides = [
            ("env == \"test\"", examples_config_v1beta1.TunableStruct(
                string_field = "test",
                number_field = 3.14,
                boolean_field = True,
            )),
        ],
    ),
)
