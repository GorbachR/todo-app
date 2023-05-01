package dev.gorbach.todoapi.repository;

import java.util.UUID;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import dev.gorbach.todoapi.dto.NoteDto;
import dev.gorbach.todoapi.entity.Note;

public interface NotesRepository extends JpaRepository<Note, UUID> {

}
