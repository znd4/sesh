package lister

func srcs(opts ListOptions) []string {
	var srcs []string
	count := 0
	if opts.Tmux {
		count++
	}
	if opts.Tmuxinator {
		count++
	}
	if opts.Config {
		count++
	}
	if opts.Zoxide {
		count++
	}
	if count == 0 {
		return []string{"tmux", "tmuxinator", "config", "zoxide"}
	}
	srcs = make([]string, count)
	i := 0
	if opts.Tmux {
		srcs[i] = "tmux"
		i++
	}
	if opts.Tmuxinator {
		srcs[i] = "tmuxinator"
		i++
	}
	if opts.Config {
		srcs[i] = "config"
		i++
	}
	if opts.Zoxide {
		srcs[i] = "zoxide"
	}
	return srcs
}
