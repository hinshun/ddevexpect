# ddevexpect

After cloning the repository, you can run `make` in the root directory of this project to execute a simple test for [ddev](https://github.com/drud/ddev)

```sh
❯ make
test
go test -v
=== RUN   TestDdevConfig
--- PASS: TestDdevConfig (0.01s)
        ddev_test.go:55: Creating a new ddev project config in the current directory (/home/edgarl/go/src
                /github.com/hinshun/ddevexpect/test)
                Once completed, your configuration will be written to /home/edgarl/go/src/github
                .com/hinshun/ddevexpect/test/.ddev/config.yaml
                Project name (test):
                The docroot is the directory from which your site is served. This is a relative
                path from your project root (/home/edgarl/go/src/github.com/hinshun/ddevexpect/t
                est)
                You may leave this value blank if your site files are in the project root
                Docroot Location (current directory):
                Found a php codebase at /home/edgarl/go/src/github.com/hinshun/ddevexpect/test.
                Project Type [php, drupal6, drupal7, drupal8, wordpress, typo3, backdrop] (php):
                Project type has no settings paths configured, so not creating settings file.
                Configuration complete. You may now run 'ddev start'.
PASS
ok      github.com/hinshun/ddevexpect   0.012s
```
