export(
    Config(
        description = "test number operators",
        default = "default",
        overrides = [
            ("version == 1", "equals"),
            ("version != 2", "not equals"),
            ("version > 3", "greater"),
            ("version >= 4", "greater or equals"),
            ("version < 5", "less"),
            ("version <= 6", "less or equals"),
            ("version in [1,3,5]", "in"),
        ],
    ),
)
