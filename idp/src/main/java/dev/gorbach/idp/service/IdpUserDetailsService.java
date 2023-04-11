package dev.gorbach.idp.service;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.Set;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;
import dev.gorbach.idp.entity.Role;
import dev.gorbach.idp.entity.User;
import dev.gorbach.idp.repository.UserRepository;

@Service
public class IdpUserDetailsService implements UserDetailsService {

    private UserRepository userRepository;

    public IdpUserDetailsService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @Override
    public UserDetails loadUserByUsername(String email) throws UsernameNotFoundException {
        User user = userRepository.findByEmail(email)
                .orElseThrow(() -> new UsernameNotFoundException("User wasn't found!"));

        return new org.springframework.security.core.userdetails.User(user.getEmail(), user.getPassword(),
                roleToGrantedAuth(user.getRoles()));
    }

    private Collection<? extends GrantedAuthority> roleToGrantedAuth(Set<Role> roles) {
        List<GrantedAuthority> grantedAuthorities = new ArrayList<>();
        for (Role role : roles) {
            grantedAuthorities.add(new SimpleGrantedAuthority(role.getName()));
        }
        return grantedAuthorities;
    }
}