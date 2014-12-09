sprinkler
====================

[WIP]

End To End Testing framework.

Get started
--------------

### Install sprinkler

OSX:

```
$ brew tap hokaccha/sprinkler
$ brew install sprinkler
```

or:

Dowonload from [releases](https://github.com/hokaccha/sprinkler/releases)

### Setup selenium

OSX:

```
$ brew install selenium-server-standalone
$ selenium-server
```

### run hello.yml

```yaml
# hello.yml
scenarios:
- name: Hello sprinkler!
  actions:
  - visit: http://www.google.com
  - assert_title: Google
  - wait_for: input[type="text"]
  - input:
      element: input[type="text"]
      value: Hello
  - submit: form[name="f"]
  - wait: 1000
  - assert_text:
      element: "#main"
      contain: Hello
```

Run it!

```
$ sprinkler hello.yml
```

See [more examples](https://github.com/hokaccha/sprinkler/tree/master/example).

License
--------------

MIT
