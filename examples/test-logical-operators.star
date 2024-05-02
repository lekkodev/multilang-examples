export(
    Config(
        description = "",
        default = "default",
        overrides = [
            ("env == \"prod\" and version == 1", "and"),
            ("env == \"test\" or version == 2", "or"),
            ("env == \"prod\" and (version == 1 or is_test == true)", "and or"),
            ("env == \"prod\" or (version == 1 and is_test == true)", "or and"),
            ("env == \"prod\" and not(version == 1 and is_test == true)", "and not"),
        ],
    ),
)
