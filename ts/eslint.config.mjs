import tseslint from 'typescript-eslint';
import lekko from "@lekko/eslint-plugin";

console.log(lekko);

export default tseslint.config(
  ...tseslint.configs.recommendedTypeChecked,
    {
    languageOptions: {
      parserOptions: {
        project: true,
        tsconfigRootDir: import.meta.dirname,
      },
    },
    files: ["lekko/*.ts"],
    plugins: {
      lekko: lekko,
    },
    rules: {
      "lekko/limitations": "error",
    },
  },
);
